# go-twirp

## Install

```
$ go get github.com/twitchtv/retool
```

Install the plugins into your project's _tools folder:

```
$ retool add github.com/golang/protobuf/protoc-gen-go master
$ retool add github.com/twitchtv/twirp/protoc-gen-twirp master
```

```
$ retool do protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./proto/HelloWorld.proto
```