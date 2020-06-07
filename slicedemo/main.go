package main

import (
	"fmt"
)

func main() {
	var intArr [6]int = [...]int{1, 2, 3, 4, 5, 6}
	// var intArr2 [3]int = [...]int{99,100,101}
	slice := intArr[:]
	copy(slice, []int{100})
	fmt.Println(slice)
	// fmt.Printf("intArr = %v", intArr)
	// fmt.Printf("slice = %v", slice[0])

	// slice[0] = 100

	// fmt.Println()
	// fmt.Printf("intArr = %v", intArr)
	// fmt.Printf("slice = %v", slice)
	// fmt.Printf("intArrPtr = %p", &intArr)
	// fmt.Printf("slicePtr = %p", &slice[0])

	// slice = append(slice, 99, 88)

	// fmt.Println()
	// fmt.Printf("intArr = %v", intArr)
	// fmt.Printf("slice = %v", slice)
	// fmt.Printf("intArrPtr = %p", &intArr)
	// fmt.Printf("slicePtr = %p", &slice[0])

	// slice2 := append(slice, slice...)

	// fmt.Println()
	// fmt.Printf("intArr = %v", intArr)
	// fmt.Printf("slice = %v", slice)
	// fmt.Printf("intArrPtr = %p", &intArr)
	// fmt.Printf("slicePtr = %p", &slice[0])
	// fmt.Printf("slice2Ptr = %p", &slice2[0])
	// slice3 := append(slice, slice...)
	// fmt.Println(intArr)
	// slice4 := append(slice3, slice...)
	// fmt.Println(intArr) 
	// slice5 := append(slice4, 3)
	// fmt.Println(intArr)
	// slice5 = append(slice5, 5,5)
	// fmt.Println(intArr)
	// fmt.Printf("intArrPtr = %p", &intArr)
	// fmt.Printf("slicePtr = %p", &slice[0])
	// fmt.Printf("slice3Ptr = %p", &slice3[0])
	// fmt.Printf("slice4Ptr = %p", &slice4[0])
	// fmt.Printf("slice5Ptr = %p", &slice5[0])
	// fmt.Println(slice5)
}
