## grpc setup step by step



## go installation

`snap install go  --classic`

## Protocol Buffer Compiler Installation

```
$ apt install -y protobuf-compiler
$ protoc --version  # Ensure compiler version is 3+
```

### Install pre-compiled binaries

```sh
$ PB_REL="https://github.com/protocolbuffers/protobuf/releases"
$ curl -LO $PB_REL/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
$ unzip protoc-3.15.8-linux-x86_64.zip -d $HOME/.local
$ export PATH="$PATH:$HOME/.local/bin"
```



## run example

```sh
$ git clone -b v1.43.0 https://github.com/grpc/grpc-go
$ cd grpc-go/examples/helloworld
$ go run greeter_server/main.go
$ go run greeter_client/main.go
Greeting: Hello world
```



其实安装按照官网一步一步来就行了，但是往往网络不太好导致安装不顺利。

我用了Azure云安装的，几分钟就完成部署了。