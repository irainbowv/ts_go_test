// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var scene sync.Map
// 	scene.Store("greece", 97)
// 	scene.Store("linda", 100)
// 	scene.Store("tst", 99)

// 	fmt.Println(scene.Load("linda"))

// 	scene.Range(func(k, v interface{}) bool {
// 		fmt.Println(k, v)
// 		return true
// 	})
// }
