package main

import (
	"fmt"
)

func main() {

	// hello world
	a := hello("world")
	fmt.Println(a)

	b, c := swap(2, 1)
	fmt.Println(b, c)

	// 無名関数を作成し、変数に代入(関数リテラル)
	tmp := func(a, b int) (int, int) {
		return b, a
	}
	fmt.Println(tmp(4, 3))

	// 即時関数
	func(msg string) {
		fmt.Println(msg)
	}("関数を指定して実行")
}

/**
 * 引数で受け取った文字列に連結させて返す(戻り値の変数を指定)
 */
func hello(msg string) (ret string) {
	ret = "hello" + msg
	return
}

/**
* 引数を逆にしてまとめて返す
 */
func swap(a, b int) (int, int) {
	return b, a
}
