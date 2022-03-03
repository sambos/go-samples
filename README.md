
## GOPATH Issues ?
[Using GoModules without worrying $GOPATH](https://medium.com/mindorks/create-projects-independent-of-gopath-using-go-modules-802260cdfb51)
* before Go 1.11 all projects had to be created inside the $GOPATH
* With the latest release of Go 1.11, Go Modules are introduced which let you create your project outside the $GOPATH and improves package management a lot.


```
$ go mod init github.com/{your_username}/{repo_name}
$ go build  -- only install/updates dependencies
$ go mod tidy -- will help you clean unused
$ go mod vendor -- for storing dependencies under /vendor 

Note: go.mod and go.sum, both needs to be added to version control.
```