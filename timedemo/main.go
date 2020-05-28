package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("3"))

	i := 1
	for {
		i++
		time.Sleep(time.Second)
		fmt.Println(i)
		if i > 10 {
			break
		}
	}
}
