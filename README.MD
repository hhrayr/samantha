# ABOUT SAMANTHA
This is a Go application used as a backend API.


# Instalation

* Install golang: https://golang.org/doc/install
* Set $GOPATH:
```
$ export GOPATH=$HOME/projects/go
```
* Clone repository into:
$GOPATH/src/github.com/hhrayr/samantha
* Install godep package manager:
```
$ go get github.com/tools/godep
```

# Building and running

* Install vscode: https://code.visualstudio.com/
* Install "Go" extension for vs from lukehoban:
* Go to source directory
```
$ cd $GOPATH/src/github.com/hhrayr/samantha
```
* Download the dependencies
```
$ godep get
```
* Open the project in vscode
```
$ code .
```
* Build and run (ctrl+shift+b) and the application is available in: http://localhost:4040/


# Development

##Get dependencies

```console
$ godep get
```

##Add new dependency
To add new package foo/bar

* Run
```
$ go get foo/bar
```
* Edit your code to import foo/bar.
* Run
```
$ godep save (or godep save ./...).
```

##Update an existing dependency
To update foo/bar package

* Run
```
$ go get -u foo/bar
```
* Run
```
$ godep update foo/bar
```