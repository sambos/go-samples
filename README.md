
## Documentation
* https://go.dev/doc/effective_go

## GOPATH Issues ?
[Using GoModules without worrying $GOPATH](https://medium.com/mindorks/create-projects-independent-of-gopath-using-go-modules-802260cdfb51)
* before Go 1.11 all projects had to be created inside the $GOPATH
* With the latest release of Go 1.11, Go Modules are introduced which let you create your project outside the $GOPATH and improves package management a lot.


### Simple setup For VsCode with Go 1.16 and above
* Install go (normally it installs under user home, if you decide to change, make sure to add to PATH)
* No need to set GOROOT or GOPATH, you may place your workspace anywhere

```
# Open empty folder in VsCode e.g "go-example"
  $ go mod init rsol.com/go-example

# create main.go
# create sub packages as you like (files placed inside should refer sub package name)

  $ go mod tidy (install any vscode recommended extensions, gopls etc)
```

Note: If you have opened multiple modules and vscode complains on imports, open module individually.

#### commands
```
$ go mod init rsol.com/root-folder
$ go build  -- only install/updates dependencies
$ go mod tidy -- will help you clean unused
$ go mod vendor -- for storing dependencies under /vendor 

Note: go.mod and go.sum, both needs to be added to version control.
```

