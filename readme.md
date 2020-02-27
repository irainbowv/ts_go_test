# Golang&&typescript 练习

## Golang的安装

下载地址：`https://golang.org/dl/`，按照对应的操作系统选择对应的安装包
安装完成后，配置环境变量

```shell
  vim ~/.bash_profile
  # insert 插入
  export GOROOT=/usr/local/go
  export GOPATH=$HOME/gowork 
  export PATH=$PATH:$GOROOT/bin:$GOPATH:$PATH
  
  # :wq 保存并退出
  source ~/.bash_profile

```

> 可以通过 `go env` 查看 go 环境变量

```go
  // 第一个hello world的demo hello.go
  package main
  import "fmt"
  func Hello() {
    ftm.Println("hello,world")
  }
  // 运行 go run hello.go
```
