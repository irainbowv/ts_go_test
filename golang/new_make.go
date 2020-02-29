// package main

// import "fmt"

// func main() {
// 	// var a = make(map[int]string)
// 	// fmt.Println(a)

// 	// for i := 1; i <= 9; i++ {
// 	// 	for j := 1; j <= i; j++ {
// 	// 		fmt.Printf("%d*%d=%d  ", j, i, i*j)
// 	// 	}
// 	// 	fmt.Println()
// 	// }

// 	c := make(chan int)
// 	go func() {
// 		c <- 1
// 		c <- 2
// 		c <- 3
// 		close(c)
// 	}()

// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }
