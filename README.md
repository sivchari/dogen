dogen 
 
"dogen" is a cli tool that provides skeletons for implementing golang layered architecture

``` tree
.
├── LICENCE
├── README.md
├── gen
│   └── dogen.go
├── go.mod
├── go.sum
├── main.go
└── templates
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
 
# DEMO
 

![スクリーンショット 2020-09-08 19 21 58](https://user-images.githubusercontent.com/55221074/92464937-decf2900-f208-11ea-9142-37f74d58c914.png)

 
# Features
 
With dogen, you can focus on your development and not get tired of creating, copying and pasting directories.
 
# Usage

 ``` command line

 dogen -m foo -t bar -d baz
 
 ```

# Note

・ Only ".tmpl" template is supported.

# Options

![スクリーンショット 2020-09-08 19 21 39](https://user-images.githubusercontent.com/55221074/92465051-0c1bd700-f209-11ea-9205-e03a97c02241.png)
 
# Author
 
* sivchari
 
# License

This software is released under the MIT License, see LICENSE.
