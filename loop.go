package main

import "fmt"

func main() {

	// 基本構文
	for i := 0; i < 3; i++ {
		if i == 0 {
			continue
		}
		fmt.Print(i)
	}

	// whileっぽく実装
	i := 0
	for i < 3 {
		fmt.Print(i)
		i++
	}

	// 無限ループ
	n := 0
	for {
		fmt.Print(n)
		if n == 5 {
			break
		}
		n++
	}
}