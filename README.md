# 



### Future work



## Build
```
go build
./rpigomtabuses
```
run the binary to see the output, recompile after updating the env!

## Getting started
uses golang!

## dev
on a 64 bit arch use
```
air
``

on 32 bit, like raspberry pi, use
```
npx nodemon --exec "go run" . --ext "go,json"  --signal SIGTERM
```


### Tooling
To use golang in nvim and with ALE only, so its fast on a headless environment, you can use gopls, gofumpt/gofmt, and goimports to handle automatic imports and opinionated formatting. 

For gopls and everything to work, make sure the gopath is set up right


```
go install golang.org/x/tools/gopls@latest
go install mvdan.cc/gofumpt@latest
go install golang.org/x/tools/cmd/goimports@latest

```
install air
```
go install github.com/air-verse/air@latest
```

``` bash
# in .profile or .**rc
export GOPATH=/usr/local/go
export PATH=$GOPATH/bin:$PATH
# add the gopath for installed libs
export PATH=$PATH:$(go env GOPATH)/bin
```
