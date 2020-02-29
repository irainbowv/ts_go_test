# Golang&&typescript 练习

## go

go 是一门 并发支持、垃圾回收的编译型系统编程语言，旨在创造一门具有静态编译语言的高性能和动态语言的高效开发之间拥有良好平衡点的一门编程语言

主要特点：

- 类型安全和内存安全
- 以非常直观和极低代价的方案实现高并发
- 高效的垃圾回收机制
- 快速编译
- 为多核计算机提供性能提升的方案
- UTF-8编码支持

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

## package/import

- Go 程序是通过package来组织的
- 只有package名为main 的包可以包含main函数
- 一个可执行程序有且仅有一个main包
- 通过import关键字导入其他非main包

- 通过const关键字进行常量的定义

- 通过在函数体外使用var关键字来进行全局变量的声明与赋值

- 通过type关键字来进行结构struct和接口interface的声明

- 通过func关键字来进行函数的声明

> 导入包之后，就可以使用格式 <PackageName>.<FuncName>来对包中的函数进行调用
>
> 如果导入包之后未调用其中的函数或者类型将会报出编译错误

当使用第三方包时，包名可能会非常接近或者相同，可以通过别名来进行区别和调用

```go
//定义时
import io "fmt"
import (
  io "fmt"
)

//使用时
io.Println("hello,world")
```

> 还可以将别名改为 `.`，在调用的时候直接写函数名，但是不推荐使用，容易混淆

## 可见性规则

- go语言中，使用大小写来决定该常量、变量、类型、接口、结构或函数 是否可以被外部包所调用
  - 根据约定，函数名首字母小写即为private （私有成员）
  - 函数名首字母大写即为public

## go基本类型

- 布尔型：bool
  - 长度：1字节
  - 取值范围：true/false，不存在与数值的隐式转换
- 整型：int/uint
  - 根据运行平台可能为32或64位
- 8位整型：int8/uint8
  - 长度：1字节
  - 取值范围：-128-127/0-255
- 字节型：byte（uint8别名）
- 复数：complex64/complex128
  - 长度：8/16字节
- 足够保存指针的32位或64位整数型：uintptr

- 其他值类型：
  - array、struct、string
- 引用类型
  - slice、map、chan

- 接口类型：inteface
- 函数类型：func

> 类型零值：零值并不等于空值，而是当变量被声明为某种类型后的默认值，通常情况下值类型的默认值为0，bool为false，string为空字符串

## 变量的类型转换

- go中不存在隐式转换，所有类型转换必须是显式声明
- 转换只能发生在两种相互兼容的类型之间

## 运算符

- go中的运算符都是从左至右结合

## 指针

go虽然保留了指针，但与其他编程语言不同的是，在go中不支持指针运算以及 "->" 运算符，而直接采用 "." 选择符来操作指针目标对象的成员

- 操作符 "&" 取变量地址，使用 "*" 通过指针间接访问目标对象
- 默认值为nil 而非 NULL

> 使用new创建数组得到的是一个指向数组的指针

## nil:空值/零值

在go语言中，布尔类型的零值为false，数值类型的零值为0，字符串类型的零值为空字符串，而指针、切片、映射、通道、函数和接口的零值则是nil

- nil是go语言中一个预定义好的标识符
- nil 标识符是不能比较的
- nil没有默认类型
- 不同类型的nil值占用的内存大小可能是不一样的

## new函数和make函数

go语言中new和make是两个内置函数，主要用来创建并分配类型的内存。new只分配内存，而make只能用于slice、map、和channel的初始化。

- new函数只接受一个参数，这个参数是一个类型，并且返回一个指向该类型内存地址的指针。同事new函数会把分配的内存置为零，也就是类型的零值。
- make只用于chan、map以及slice的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以没有必要返回他们的指针了。

## for range 键值循环

for range 可以遍历数组、切片、字符串、map 及通道（channel），for range 语法上类似于其它语言中的 foreach 语句

```go
  // for key,val := range x {
  //   ...
  // }
```

val始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对他所做的任何修改都不会影响到集合中原有的值。

- go语言的for包含初始化语句、条件表达式、结束语句，这三个部分均可缺省
- 在需要时，可以使用匿名变量对for range的变量进行选取

## goto break continue

- Go语言中 goto 语句通过标签进行代码间的无条件跳转，同时 goto 语句在快速跳出循环、避免重复退出上也有一定的帮助，使用 goto 语句能简化一些代码的实现过程。
- Go语言中 break 语句可以结束 for、switch 和 select 的代码块，另外 break 语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的 for、switch 和 select 的代码块上。
- Go语言中 continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用，在 continue 语句后添加标签时，表示开始标签对应的循环

## 函数

- 在函数中，实参通过值传递的方式进行传递，因此函数的形参是实参的拷贝，对形参进行修改不会影响实参，但是，如果实参包括引用类型，如指针、slice(切片)、map、function、channel 等类型，实参可能会由于函数的间接引用被修改
- Go语言支持多返回值，多返回值能方便地获得函数执行后的多个返回参数，Go语言经常使用多返回值中的最后一个返回参数返回函数执行中可能发生的错误

```go
  conn,err := connectToNetwork()
```
