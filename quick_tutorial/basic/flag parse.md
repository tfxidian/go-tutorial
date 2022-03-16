# flag parse

为了看的更清楚，这里写一个极简的例子：
```go
package main

import (
	"flag"
	"fmt"
)

var cliName = flag.String("name", "John", "Input Your Name")

func main() {
	flag.Parse()
	fmt.Println("name=", *cliName)

}

```


 定义命令行参数对应的变量
```go
var cliName = flag.String("name", "nick", "Input Your Name")
```

定义完命令行参数对应的变量后，通过 flag.Parse() 函数解析命令行参数。
```go
flag.Parse()
```

使用的时候就用 `go run mian.go --name=ftang `即可



