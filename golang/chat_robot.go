// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	// 准备从标准流中读取输入数据
// 	inputReader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Please input your name:")
// 	// 读取数据直到碰到\n为止
// 	input, err := inputReader.ReadString('\n')
// 	if err != nil {
// 		fmt.Printf("An error  occurred : %s\n", err)
// 		// 异常退出
// 		os.Exit(1)

// 	} else {
// 		name := input[:len(input)-1]
// 		fmt.Printf("Hello, %s ! What can  I do for you ?\n", name)
// 	}

// 	for {
// 		input, err = inputReader.ReadString('\n')
// 		if err != nil {
// 			fmt.Printf("An error  occurred : %s\n", err)
// 			continue
// 		}
// 		input = input[:len(input)-1]
// 		input = strings.ToLower(input)
// 		fmt.Println("input", input)
// 		switch input {
// 		case "":
// 			continue
// 		case "nothing", "bye":
// 			fmt.Println("bye")
// 			os.Exit(0)
// 		default:
// 			fmt.Println("Sorry , I didn't catch you.")
// 		}
// 	}
// }
