package gen

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	version = "ca-dojo version"
	// ShowVersion returns when set version flag.
	ShowVersion = errors.New("show version")
)

// cli options
type Dogen struct {
	params    Params
	dir       string
	template  string
	outStream io.Writer
	errStream io.Writer
}

type Params struct {
	Dist  string
	Model string
	Root  string
}

func fill(args []string, outStream, errStream io.Writer) (*Dogen, error) {
	var v bool
	var m, t, d string
	// dogen
	toolName := args[0]
	flags := flag.NewFlagSet(toolName, flag.ContinueOnError)
	// 標準エラーへ出力先を変更
	flags.SetOutput(errStream)

	versionDescribe := "print dogen version information"
	flags.BoolVar(&v, "version", false, versionDescribe)
	flags.BoolVar(&v, "v", false, versionDescribe)

	modelDescribe := "model name"
	flags.StringVar(&m, "model", "", modelDescribe)
	flags.StringVar(&m, "m", "", modelDescribe)

	tmplDescribe := "templates directory"
	flags.StringVar(&t, "tmpl", "", tmplDescribe)
	// 指定なしで./pure指定
	flags.StringVar(&t, "t", "pure", tmplDescribe)

	dirDescribe := "output directory"
	flags.StringVar(&d, "dir", "", dirDescribe)
	// 指定なしでcurrent dirに生成
	flags.StringVar(&d, "d", "pkg", dirDescribe)

	// dogen optionのパース
	if err := flags.Parse(args[1:]); err != nil {
		return nil, err
	}

	// version option
	if v {
		fmt.Fprintf(outStream, "%s version %s\n", toolName, version)
		return nil, ShowVersion
	}

	// cli optionでモデル名を指定していない場合エラー
	// ディレクトリとテンプレートは指定なしでもデフォルトで通る
	if len(m) == 0 {
		msg := "please enter options"
		return nil, fmt.Errorf(msg)
	}

	// model名lowerケース
	m = strings.ToLower(m)

	dogenArgs := flags.Args()
	if len(dogenArgs) > 0 {
		msg := "please enter only one arg"
		return nil, fmt.Errorf(msg)
	}

	// templateとdirectoryの絶対パスを取得
	t, err := filepath.Abs(t)
	if err != nil {
		return nil, err
	}
	dir, err := filepath.Abs(d)
	if err != nil {
		return nil, err
	}

	// カレントディレクトリの名前取得
	p, _ := os.Getwd()
	r := strings.Split(p, "/")
	root := r[len(r)-1]

	dogen := &Dogen{
		params: Params{
			Dist:  d,
			Model: m,
			Root:  root,
		},
		dir:       dir,
		template:  t,
		outStream: outStream,
		errStream: errStream,
	}
	return dogen, nil
}

func (dogen *Dogen) run() error {
	// templateの各ファイルに対してwalk関数を実行
	return filepath.Walk(dogen.template, dogen.walk)
}

var fmap = template.FuncMap{
	"title": strings.Title,
}

// make file name with action and model.
func (dogen *Dogen) buildFileName(base string) string {
	//tmplをgoに変換し、
	base = strings.Replace(base, ".tmpl", ".go", 1)
	return strings.ToLower(strings.Join([]string{dogen.params.Model, base}, "_"))
}

func (dogen *Dogen) walk(path string, info os.FileInfo, err error) error {
	p, err := filepath.Rel(dogen.template, path)
	if err != nil {
		return err
	}

	fp := filepath.Join(dogen.dir, p)

	// walkした先がdirなら同様の物をmkdir
	if info.IsDir() {
		if err := os.MkdirAll(fp, 0777); err != nil {
			return err
		}
		return nil
	}

	// tmpl拡張子以外はerr
	if filepath.Ext(path) != ".tmpl" {
		return nil
	}

	dn, fn := filepath.Split(fp)
	// tmplをgoに拡張子変更
	sp := filepath.Join(dn, dogen.buildFileName(fn))

	buf := bytes.Buffer{}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// pathからtemlateを解析し、option引数を基にparse
	dtmpl := string(b)
	if err := template.Must(template.New(sp).Funcs(fmap).Parse(dtmpl)).Execute(&buf, dogen.params); err != nil {
		return err
	}

	// 作成した*.goへgo fmtを実行しファイル生成
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
}

// Run is entry point.
func Run(args []string, outStream, errStream io.Writer) error {
	dogen, err := fill(args, outStream, errStream)
	if err != nil {
		return err
	}
	return dogen.run()
}
