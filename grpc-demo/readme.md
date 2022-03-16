

本文完全参考：[最简单的 gRPC 教程— 1 初识 gRPC](https://zhuanlan.zhihu.com/p/359968500)
目录结构：
```
grpc-demo/
├── client
│   └── main.go
├── go.mod
├── go.sum
├── product
│   ├── ProductInfo.pb.go
│   └── ProductInfo.proto
└── server
    └── main.go
```

需要注意，我一直以为go mod init后面跟一个随意的名字就好了，今天编这个项目才知道，是要跟包名：
 `go mod init grpc-demo`
否则，在调用的时候可能出现
```
"package XXX is not in GOROOT" when building a Go project
```
这种错误。