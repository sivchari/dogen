package dogen

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkRunSynchronous(b *testing.B) {
	d := dogen{
		params: params{
			Dist:  "once_pkg",
			Model: "user",
			Root:  "dogen",
		},
		dir:      "/Users/pc-351/workspace/go/dogen/cmd/once_pkg",
		template: "/Users/pc-351/workspace/go/dogen/cmd/pure",
		mu:       sync.Mutex{},
	}
	//最初に長さを決める
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		// 都度append
		err := d.run()
		if err != nil {
			log.Print(err)
		}
	}
}

func BenchmarkRunAsynchronousGoroutine(b *testing.B) {
	d := dogen{
		params: params{
			Dist:  "routine_pkg",
			Model: "user",
			Root:  "dogen",
		},
		dir:      "/Users/pc-351/workspace/go/dogen/cmd/routine_pkg",
		template: "/Users/pc-351/workspace/go/dogen/cmd/pure",
		mu:       sync.Mutex{},
	}
	//最初に長さを決める
	b.ResetTimer()

	runtime.GOMAXPROCS(runtime.NumCPU())
	eg, ctx := errgroup.WithContext(context.Background())
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		for {
			eg.Go(func() error {
				dogen := d
				select {
				case <-ctx.Done():
					break
				default:
					dogen.mu.Lock()
					defer dogen.mu.Unlock()
					if err := dogen.run(); err != nil {
						return errors.New("[ERROR] Fatal Error: " + err.Error())
					}
					return nil
				}
				return nil
			})
			goto Done
		}
	Done:
		if err := eg.Wait(); err != nil {
			log.Print(err)
		}
	}
}
