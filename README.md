# dogen

## `dogen` is a cli tool that provides skeletons for implementing golang layered architecture

# Features
With dogen, you can focus on your development and not get tired of creating, copying and pasting directories.

# Install
```shell
git clone git@github.com:sivchari/dogen.git
```

If you use default setting, please following below
```shell
go get github.com/sivchari/dogen
```

# Tree
```
.
├── LICENCE
├── README.md
├── dogen
│   └── dogen.go
├── go.mod
├── go.sum
├── main.go
├── statik
│   └── statik.go
└── tmpl
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

# Usage
```sh
cp .env.example .env
```

## Set your go mod path to .env
```env
DOGEN_PKG=github.com/xxx/yyy
```

## Generate skelton
```sh
dogen -m user
```

```sh
pkg/
├── domain
│   ├── model
│   │   └── user_model.go
│   └── repository
│       ├── mock_repository
│       │   └── mock_user_repository.go
│       └── user_repository.go
├── infrastructure
│   └── mysql
│       └── repoimpl
│           ├── mock_repoimpl
│           │   └── mock_user_repoimpl.go
│           └── user_repoimpl.go
├── interfaces
│   └── handler
│       ├── mock_handler
│       │   └── mock_user_handler.go
│       └── user_handler.go
└── usecase
    ├── mock_usecase
    │   └── mock_user_usecase.go
    └── user_usecase.go
```

## File content (user_handler.go)
```go
//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"github.com/labstack/echo/v4"

	userU "sample/pkg/usecase"
)

type UserHandler interface {
	HandleSelect(c echo.Context) echo.HandlerFunc
	HandleInsert(c echo.Context) echo.HandlerFunc
	HandleUpdate(c echo.Context) echo.HandlerFunc
	HandleDelete(c echo.Context) echo.HandlerFunc
}

type userHandler struct {
	userUseCase userU.UserUseCase
}

// NewUserHandler
func NewUserHandler(userU userU.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: userU,
	}
}

// HandleSelect
func (userH *userHandler) HandleSelect(c echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("do something")
		return nil
	}
}

// HandleInsert
func (userH *userHandler) HandleInsert(c echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("do something")
		return nil
	}
}

// HandleUpdate
func (userH *userHandler) HandleUpdate(c echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("do something")
		return nil
	}
}

// HandleDelete
func (userH *userHandler) HandleDelete(c echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("do something")
		return nil
	}
}

// UserRequest
type UserRequest struct {
	// Need to implement field
}

// UserResponse
type UserResponse struct {
	// Need to implement field
}
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
  -g string
    	generate type (echo or pure) (default "echo")
  -gen string
    	generate type (echo or pure)
  -m string
    	model name
  -model string
    	model name
  -v	print dogen version information
  -version
    	print dogen version information
```

#benchmark
```
BenchmarkRunSynchronous-4                     87          13434437 ns/op          357120 B/op       6183 allocs/op
BenchmarkRunAsynchronousGoroutine-4       147633              8387 ns/op              42 B/op          2 allocs/op

```

# Author

## sivchari
## [Twitter](https://twitter.com/sivchari)

# License
This software is released under the MIT License, see LICENSE.
