package main

import "fmt"

func main() {

	// スライスを作成
	e := []int{1, 3, 5}

	// 要素を追加
	e = append(e, 7, 9)
	fmt.Println(e)

	// eの要素数のスライスを作成
	f := make([]int, len(e))

	// eスライスをコピーし、コピーした要素数をgに代入
	g := copy(f, e)
	fmt.Println(f)
	fmt.Println(g)
}
