# TeamGandalf
This project is in Go.

To understand initial set up see https://golang.org/doc/code.html

Clone this repo into your Go work src repository.

Add code to main and build off that.

Dependencies:

https://github.com/sjwhitworth/golearn/wiki/Installation

https://github.com/satori/go.uuid

# Building the project

The following commands should get the source files and all the dependencies.
There is an error about non-local imports that can be safely ignored.


```sh
go get github.com/AKeizer/TeamGandalf
cd $GOPATH/src/github.com/AKeizer/TeamGandalf
go get ./...
go install -a main.go
```

# Website

Running the code with :

./main -web

Opens a web server that is a demonstration of the ML algorithm based on the trained
data we have generated.

main page is http://localhost:8080/main.html
