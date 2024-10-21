# Cryptopals workthrough

This is a repo that I created while working through the [cryptopals](https://cryptopals.com) challenges.  Decided to use golang since it's become my choice for doing ctfs and its a bit of a new language that is very close to C.


### Future work



## Build

``` bash
# go build
./rpigomtabuses
```

run the binary to see the output, recompile after updating the env!

## Getting started
uses golang!

## dev
on a 64 bit arch use

``` bash
air
```

on 32 bit, like raspberry pi, use
``` bash
npx nodemon --exec "go run" . --ext "go,json"  --signal SIGTERM
```


### Tooling
To use golang in nvim and with ALE only, so its fast on a headless environment, you can use gopls, gofumpt/gofmt, and goimports to handle automatic imports and opinionated formatting. 

For gopls and everything to work, make sure the gopath is set up right


``` bash
go install golang.org/x/tools/gopls@latest
go install mvdan.cc/gofumpt@latest
go install golang.org/x/tools/cmd/goimports@latest

```
install air
``` bash
go install github.com/air-verse/air@latest
```

``` bash
# in .profile or .**rc
export GOPATH=/usr/local/go
export PATH=$GOPATH/bin:$PATH
# add the gopath for installed libs
export PATH=$PATH:$(go env GOPATH)/bin
```
