package main

import (
	"fmt"
	"strings"
)

// func add() func(int) int {
// 	n := 20
// 	return func(x int) int {
// 		n += x
// 		return n
// 	}
// }
func makeSuffix(s1 string) func(string) string {
	return func(s2 string) string {
		if !strings.HasSuffix(s2, s1) {
			return s2 + s1
		}
		return s2
	}
}

func main() {
	// a := add()
	// fmt.Println(a(10))
	// fmt.Println(a(10))
	// fmt.Println(a(10))
	a := makeSuffix(".jpg")
	fmt.Println(a("test1.jpg"))
	fmt.Println(a("test2"))
	fmt.Println(a("test3"))
}
