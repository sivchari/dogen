dogen 
 
"dogen" is a cli tool that provides skeletons for implementing golang layered architecture

# Features
With dogen, you can focus on your development and not get tired of creating, copying and pasting directories.

# Install
```shell
git clone git@github.com:sivchari/dogen.git
```

edit dogen/dogen.go
```go
Engine = "pure" OR Engine = "echo" (Default is pure.)
```
```shell
statik -src=pure OR statik -src=echo (Please match the engine you selected.)
go build -o dogen
```

If you use default setting, please following below
```shell
go get github.com/sivchari/dogen
```

# Usage
 ``` command line
 dogen -m XXX
 ```

# Example
```
.
├── LICENCE
├── Makefile
├── README.md
├── dogen
│   ├── dogen.go
│   └── dogen_test.go
├── echo
│   ├── domain
│   │   ├── model
│   │   │   └── model.tmpl
│   │   └── repository
│   │       └── repository.tmpl
│   ├── infrastructure
│   │   └── mysql
│   │       └── repoimpl
│   │           └── repoimpl.tmpl
│   ├── interfaces
│   │   └── handler
│   │       └── handler.tmpl
│   └── usecase
│       └── usecase.tmpl
├── example
│   ├── domain
│   │   ├── model
│   │   │   └── user_model.go
│   │   └── repository
│   │       └── user_repository.go
│   ├── infrastructure
│   │   └── mysql
│   │       └── repoimpl
│   │           └── user_repoimpl.go
│   ├── interfaces
│   │   └── handler
│   │       └── user_handler.go
│   └── usecase
│       └── user_usecase.go
├── go.mod
├── go.sum
├── main.go
├── pure
│   ├── domain
│   │   ├── model
│   │   │   └── model.tmpl
│   │   └── repository
│   │       └── repository.tmpl
│   ├── infrastructure
│   │   └── mysql
│   │       └── repoimpl
│   │           └── repoimpl.tmpl
│   ├── interfaces
│   │   └── handler
│   │       └── handler.tmpl
│   └── usecase
│       └── usecase.tmpl
└── statik
   └── statik.go
```

# Note
・ Only ".tmpl" template is supported.

# Options
```
Usage of dogen:
  -d string
        output directory (default "pkg")
  -dir string
        output directory
  -m string
        model name
  -model string
        model name
  -v    print dogen version information
  -version
        print dogen version information
ʕ◔ϖ◔ʔ 🍭 
```

#benchmark
```
BenchmarkRunSynchronous-4                     87          13434437 ns/op          357120 B/op       6183 allocs/op
BenchmarkRunAsynchronousGoroutine-4       147633              8387 ns/op              42 B/op          2 allocs/op

```

 
# Author

##sivchari
## [Twitter](https://twitter.com/sivchari)
 
# License
This software is released under the MIT License, see LICENSE.
