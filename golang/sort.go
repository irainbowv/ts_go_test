// package main

// import (
// 	"fmt"
// 	"strings"
// )

// var original = []string{
// 	"Nonmetals",
// 	"    Hydrogen",
// 	"    Carbon",
// 	"    Nitrogen",
// 	"    Oxygen",
// 	"Inner Transitionals",
// 	"    Lanthanides",
// 	"        Europium",
// 	"        Cerium",
// 	"    Actinides",
// 	"        Uranium",
// 	"        Plutonium",
// 	"        Curium",
// 	"Alkali Metals",
// 	"    Lithium",
// 	"    Sodium",
// 	"    Potassium",
// }

// func main() {
// 	a, b := computeIndent(original)
// 	fmt.Println(len(a), b)
// }
// func computeIndent(slice []string) (string, int) {
// 	for _, item := range slice {
// 		if len(item) > 0 && (item[0] == ' ' || item[0] == '\t') {
// 			whitespace := rune(item[0])
// 			fmt.Println(len(item[1:]))
// 			for i, char := range item[1:] {
// 				if char != whitespace {
// 					i++
// 					return strings.Repeat(string(whitespace), i), i
// 				}
// 			}
// 		}
// 	}
// 	return "", 0
// }
