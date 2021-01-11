dogen 
 
"dogen" is a cli tool that provides skeletons for implementing golang layered architecture

``` tree
.
├── LICENCE
├── README.md
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
├── dogen
│   └── dogen.go
├── go.mod
├── go.sum
├── main.go
└── pure
    ├── domain
    │   ├── model
    │   │   └── model.tmpl
    │   └── repository
    │       └── repository.tmpl
    ├── infrastructure
    │   └── mysql
    │       └── repoimpl
    │           └── repoimpl.tmpl
    ├── interfaces
    │   └── handler
    │       └── handler.tmpl
    └── usecase
        └── usecase.tmpl
```
 
 
# Features
With dogen, you can focus on your development and not get tired of creating, copying and pasting directories.

# Install
```
1. 
go get -u https://github.com/sivchari/dogen

2.
go get -u github.com/golang/mock/gomock
go get -u github.com/golang/mock/mockgen
go get -u github.com/rakyll/statik

OR

go get -u all

3.
edit dogen/dogen.go

Engine = "pure" OR Engine = "echo" (Default is pure.)

4.
statik -src=pure OR statik -src=echo (Please match the engine you selected.)

5. 
go build -o dogen
```

# Usage
 ``` command line
 dogen -m XXX
 ```

# Example
```
dogen -u model -d pkg  

.
├── dogen
└── pkg
    ├── domain
    │   ├── model
    │   │   └── user_model.go
    │   └── repository
    │       └── user_repository.go
    ├── infrastructure
    │   └── mysql
    │       └── repoimpl
    │           └── user_repoimpl.go
    ├── interfaces
    │   └── handler
    │       └── user_handler.go
    └── usecase
        └── user_usecase.go

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
### CPU : 4
```
BenchmarkRunSynchronous-4                   336                   3894149 ns/op          361212 B/op       6177 allocs/op
BenchmarkRunAsynchronousGoroutine-4         1292509               916 ns/op              208 B/op          2 allocs/op
```

 
# Author

##sivchari
## [Twitter](https://twitter.com/sivchari)
 
# License
This software is released under the MIT License, see LICENSE.
