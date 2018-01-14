package main

import "fmt"

func main() {
	// 配列を宣言(0初期化)
	var a [5]int
	fmt.Println(a)

	// 配列を宣言して代入
	b := [3]int{1, 3, 5}
	fmt.Println(b[2])

	// [...]で要素の大きさを省略
	c := [...]int{2, 4, 6}
	fmt.Println(len(c))
}
