package main

import "fmt"

var (
	Fun = func(a int, b int) int {
		return a * b
	}
)

func main() {
	func(a int, b int) {
		fmt.Println(a + b)
	}(4, 5)

	fmt.Println(Fun(4, 5))
}
