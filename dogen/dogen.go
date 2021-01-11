package dogen

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/rakyll/statik/fs"
	_ "github.com/sivchari/dogen.git/statik"
	"go/format"
	"golang.org/x/sync/errgroup"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"text/template"
)

// dogen Version
const Version = "v1.0"

// ShowVersion returns when set version flag.
var ShowVersion = errors.New("show version")

// Select Template Engine
const Engine = "pure"

// Extenstion
const Extension = ".tmpl"

// params command args
type params struct {
	Dist  string
	Model string
	Root  string
}

// dogen  cli options
type dogen struct {
	params   params
	dir      string
	template string
	mu       sync.Mutex
}

// Run mapping and fs.Walk
func Run(args []string, outStream io.Writer, errStream io.Writer) error {
	d, m, err := setFlags(args, outStream, errStream)
	if err != nil {
		return err
	}
	dogen, err := fill(d, m)
	if err != nil {
		return err
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	eg, ctx := errgroup.WithContext(context.Background())
	for {
		eg.Go(func() error {
			dogen := dogen
			select {
			case <-ctx.Done():
				return nil
			default:
				dogen.mu.Lock()
				defer dogen.mu.Unlock()
				if err := dogen.run(); err != nil {
					return errors.New("[ERROR] Error: " + err.Error())
				}
				return nil
			}
		})
		goto Done
	}
Done:
	if err := eg.Wait(); err != nil {
		log.Print(err)
	}
	return nil
}

func setFlags(args []string, outStream, errStream io.Writer) (string, string, error) {

	var (
		v    bool
		m, d string
	)

	toolName := args[0]

	flags := flag.NewFlagSet(toolName, flag.ContinueOnError)

	flags.SetOutput(errStream)

	versionDescribe := "print dogen version information"
	flags.BoolVar(&v, "version", false, versionDescribe)
	flags.BoolVar(&v, "v", false, versionDescribe)

	modelDescribe := "model name"
	flags.StringVar(&m, "model", "", modelDescribe)
	flags.StringVar(&m, "m", "", modelDescribe)

	dirDescribe := "output directory"
	flags.StringVar(&d, "dir", "", dirDescribe)
	flags.StringVar(&d, "d", "pkg", dirDescribe)

	if err := flags.Parse(args[1:]); err != nil {
		return "", "", err
	}

	if v {
		fmt.Fprintf(outStream, "%s version %s\n", toolName, Version)
		return "", "", ShowVersion
	}

	if len(m) == 0 {
		msg := "please enter options"
		return "", "", fmt.Errorf(msg)
	}

	dogenArgs := flags.Args()
	if len(dogenArgs) > 0 {
		msg := "please enter only one arg"
		return "", "", fmt.Errorf(msg)
	}

	return d, m, nil
}

// fill mapping commands
func fill(d string, m string) (*dogen, error) {

	lm := strings.ToLower(m)

	t, err := filepath.Abs(Engine)
	if err != nil {
		return nil, err
	}

	dir, err := filepath.Abs(d)
	if err != nil {
		return nil, err
	}

	p, _ := os.Getwd()
	r := strings.Split(p, "/")
	root := r[len(r)-1]

	return &dogen{
		params: params{
			Dist:  d,
			Model: lm,
			Root:  root,
		},
		dir:      dir,
		template: t,
	}, nil
}

func (dogen *dogen) fmap() template.FuncMap {
	return template.FuncMap{
		"title": strings.Title,
	}
}

// make file name with action and model.
func (dogen *dogen) buildFileName(base string) string {
	base = strings.Replace(base, ".tmpl", ".go", 1)
	return strings.ToLower(strings.Join([]string{dogen.params.Model, base}, "_"))
}

func (dogen *dogen) run() error {

	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	err = fs.Walk(statikFS, "/", func(path string, info os.FileInfo, err error) error {
		p, err := filepath.Rel(dogen.template, path)
		if err != nil {
			return err
		}

		fp := filepath.Join(dogen.dir, p)

		if info.IsDir() {
			if err := os.MkdirAll(dogen.dir+fp, 0777); err != nil {
				return err
			}
			return nil
		}

		if filepath.Ext(path) != Extension {
			return errors.New("not allowed to parse")
		}

		dn, fn := filepath.Split(dogen.dir + fp)

		sp := filepath.Join(dn, dogen.buildFileName(fn))

		file, err := statikFS.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		b, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		buf := bytes.Buffer{}
		fmap := dogen.fmap()

		if err := template.Must(
			template.New(sp).
				Funcs(fmap).
				Parse(string(b))).
			Execute(&buf, dogen.params); err != nil {
			return err
		}

		codes, err := format.Source(buf.Bytes())
		if err != nil {
			return err
		}

		f, err := os.Create(sp)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err = f.Write(codes); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
