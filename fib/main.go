package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 1

	fmt.Println(a)

	for i := 0; i < 60; i++ {
		fmt.Println(b)
		tmp := a
		a = b
		b = tmp + a
	}
}
