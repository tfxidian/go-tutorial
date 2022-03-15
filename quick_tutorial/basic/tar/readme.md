## go 遍历目录
code:
```go
func GetFiles()  {
	files, err := ioutil.ReadDir(".")
	if err != nil{
		log.Fatal(err)
	}

	for _, file := range files{
		fmt.Println(file.Name())
	}
}
```


执行`go run test.go`
```shell script
test begin
go.mod
test.go

```

