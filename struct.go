package main

import "fmt"

// 構造体の宣言
type user struct {
	first_name string
	last_name  string
	score      int
}

func main() {

	u := new(user) // アドレスが返る

	// 代入方法1
	u.first_name = "Taro"

	// 代入方法2
	(*u).last_name = "Yamada"

	// 代入方法3 フィールド順
	uu := user{"Taro", "Yamada", 100}

	// 代入方法4 key指定
	uuu := user{first_name: "Taro", last_name: "Yamada", score: 100}

	fmt.Println(u)
	fmt.Println(uu)
	fmt.Println(uuu)
}
