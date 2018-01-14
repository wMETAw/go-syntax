package main

import "fmt"

func main() {

	// スライスを作成
	slice := []int{2, 3, 8}

	// slice文のループ
	for idx, v := range slice {

		// rangeは複数の値を返す
		fmt.Println(idx, v)
	}

	// 値を破棄したければ、ブランク修飾子「 _ 」を使用
	for _, v := range slice {
		fmt.Println(v)
	}

	// mapの場合は key valueが返る
	m := map[string]string{"first_name": "taro", "last_name": "yamada"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
