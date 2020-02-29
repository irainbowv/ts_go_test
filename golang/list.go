// package main

// import (
// 	"container/list"
// 	"fmt"
// )

// func main() {
// 	// l := list.New()
// 	var l list.List
// 	l.PushBack("fsd")
// 	l.PushFront(32)
// 	// 尾部添加后保存元素句柄
// 	element := l.PushBack("fist")
// 	// 在fist之后添加high
// 	l.InsertAfter("high", element)
// 	// 在fist之前添加noon
// 	l.InsertBefore("noon", element)
// 	// 使用
// 	// l.Remove(element)
// 	fmt.Println(element)

// 	for i := l.Front(); i != nil; i = i.Next() {
// 		fmt.Println(i.Value)
// 	}
// }
