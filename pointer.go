package main

import (
	"fmt"
)

func main() {

	a := 5

	// int型のアドレス宣言
	var pa *int
	pa = &a

	fmt.Println(*pa)
	fmt.Println(pa)
}
