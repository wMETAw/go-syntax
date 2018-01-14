package main

import "fmt"

func main() {

	// 要素数を指定せず、配列を作成(コンパイラが数える)
	c := [...]int{1, 3, 5, 7, 9}

	// ２番目から(4-1)番目をスライス
	d := c[2:4]
	fmt.Println(d)
	fmt.Println(c)

	// スライスで生成した配列に値を代入 ※元の配列の参照なのでc[4]は12となる
	d[1] = 12
	fmt.Println(c)

	// len() 配列の長さ
	// cap() 配列の先頭から切り出せる最大数
	fmt.Println(d, len(d), cap(d))
}
