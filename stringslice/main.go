package main

import (
	"fmt"
)

func main() {
	str := "abccd@hello"
	slice := str[:5]
	fmt.Println(slice)
	slice2 := []rune(str)
	slice2[1] = '龙'
	str = string(slice2)
	println(str)
}