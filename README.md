
# GoLang

Guide to setup `golang` development environment on MacOS without admin privileges

### Install

#### Directory structure

HOME = ~/

> I am organizing Directory structure as:

```bash
mkdir $HOME/go  # this is $GOPATH
mkdir $HOME/do/Developer/Apps # this os $GOROOT
mkdir $HOME/do/Developer/Apps/go # this where I download `go` and added to PATH
mkdir $HOME/do/Developer/Work # this is top level Directory for all my code/projects
mkdir $HOME/do/Developer/Work/go # this where I keep my go projects
mkdir $HOME/do/Developer/Work/java # this where I keep my java projects
```

#### Download

```bash
mkdir  -p ~/Developer/Apps/go
# download and unpack
curl -s https://dl.google.com/go/go1.12.5.darwin-amd64.tar.gz | tar xvz - -C ~/Developer/Apps
# or in silent mode
curl -s https://dl.google.com/go/go1.12.5.darwin-amd64.tar.gz | tar zxf - -C ~/Developer/Apps
```

#### Workspace

> Setup Go global Workspace. this is your `GOPATH`

```bash
mkdir -p ~/go
```

#### Environment

> export `GOROOT`, `GOPATH` and `PATH` in `.bash_profile` or similar

```bash
export GOROOT=$HOME/Developer/Apps/go
export GOPATH=$HOME/go

export PATH=${PATH}:${GOROOT}/bin:$GOPATH/bin

export PATH
```

---

### Project

#### Your first project

> create a project directory outside of your GOPATH to use Go modules:
> you go workspace is `~/Developer/Work/go`

```bash
mkdir  -p ~/Developer/Work/go/hello
```

#### Initialize module

> this will create `go.mod` file

```bash
cd ~/Developer/Work/go/hello

go mod init github.com/xmlking/hello
```

#### Write your code

```bash
$ cat <<EOF > hello.go
package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
EOF
```

#### Docs

```bash
# read docs
go doc rsc.io/quote
go help build
```

#### Format

> format your source code with `prettier` or `fmt`

```bash
go fmt hello.go
```

#### Build and Run

```bash
go build
# run
./hello
#  ./... pattern matches all the packages within the current module.
go build ./...
```

> Note: The go.mod file was updated to include explicit versions for your dependencies
> In contrast, `go build` and `go test` will not remove dependencies from `go.mod` that are no longer required and only update `go.mod` based on the current build invocation's tags/OS/architecture.

#### Test

```bash
# test all the go files within the current directory
go test
# test all the packages within the current module
go test ./...
# tests for your module plus the tests for all direct and indirect dependencies to check for incompatibilities
go test all
```

#### Install App

> this will copy binary into `$GOPATH/bin` i.e.,  `~/go/bin`

```bash
go install
```

#### Release

> recomdned steps for releasing modules

```bash
# to possibly prune any extraneous requirements
go mod tidy
#  test and  validate your module
go test all
```

Ensure your `go.sum` file is committed along with your `go.mod` file.

A new module version may be published by pushing a tag to the repository that contains the module source code.

##### Releasing Modules (v2 or Higher)

`TODO`

#### Daily Workflow

> Note there was no `go get` required in the example above.

Your typical day-to-day workflow can be:

* Add import statements to your `.go` code as needed.
* Standard commands like `go build` or `go test` will automatically add new dependencies as needed to satisfy imports (updating `go.mod` and downloading the new dependencies).
* When needed, more specific versions of dependencies can be chosen with commands such as `go get foo@v1.2.3`, `go get foo@master`, `go get foo@e3702bed2`, or by editing go.mod directly.

#### Install Additional Dev Tools when needed

> gopls, gocode, gopkgs, go-outline, go-symbols, guru, gorename, dlv, gocode-gomod, godef, goreturns, golint

```bash
go get -u github.com/golang/lint/golint
# Then, run your static analysis tools (at least, golint)
golint app.go
go get -u -v golang.org/x/tools/cmd/gopls

```

---

### Go Commands

```bash
go env
#  View final versions that will be used in a build
go list -m all
#  View available minor and patch upgrades for all direct and indirect dependencies
go list -u -m all
#  List the available tagged versions of that module
go list -m -versions rsc.io/sampler
#  Update all direct and indirect dependencies to latest minor or patch upgrades
go get -u  # to use the latest minor or patch releases
go get -u=patch # to use the latest patch releases
#  Updates to the latest version of `foo` plus direct and indirect dependencies of `foo`
go get foo  # is equivalent to `go get foo@latest`
go get -u foo # solely gets the latest version of `foo`
#  Build or test all packages in the module when run from the module root directory
go build ./...
go test ./...
# lists all the commands go build invokes.
go build -x

go test -v context
go test -race
#  Prune any no-longer-needed dependencies from go.mod and add any dependencies needed
go mod tidy
#  If you are curious why a particular module is showing up in your go.mod, then run
go mod why -m <module>
#  Optional step to create a vendor directory (use special case only)
#  Go modules ignores the vendor/ directory by default. The idea is to eventually do away with vendoring
go mod vendor
```

---

### Points

* A repository contains one or more Go modules.
* Each module contains one or more Go packages.
* Each package consists of one or more Go source files in a single directory.
* A module is defined by a tree of Go source files with a `go.mod` file in the tree's root directory.
* The import paths for all packages in a module share the module path as a common prefix.
* The module path and the relative path from the `go.mod` to a package's directory together determine a package's import path.

---

### FAQ

1. What `go build` do?

What the go command does depends on whether we run it for a "normal" package or for the special  "main" package.

**For packages**
    * `go build`   builds your package then discards the results.
    * `go install` builds then installs the package in your $GOPATH/pkg directory.

**For commands (package main)**
    * `go build`   builds the command and leaves the result in the current working directory.
    * `go install` builds the command in a temporary directory then moves it to $GOPATH/bin.

Basically you can use `go build` as a check that the packages can be built (along with their dependencies) while `go install` also (permanently) installs the results in the proper folders of your  $GOPATH.

It is also worth noting that starting with Go 1.5 `go install` also removes executables created by go build

---

### Reference

1. <https://github.com/golang/go/wiki/Modules>
2. <https://blog.golang.org/using-go-modules>
3. Project Layout
    1. <https://github.com/golang-standards/project-layout>
    2. <https://code.fb.com/security/service-encryption/>
4. <https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-daebb36a97e9>
