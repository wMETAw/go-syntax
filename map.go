package main

import (
	"fmt"
)

func main() {

	// map[keyの型]valueの型
	m := make(map[string]string)
	m["first_name"] = "taro"
	m["last_name"] = "yamada"
	fmt.Println(m)

	// 宣言と代入
	i := map[string]string{"first_name": "taro", "last_name": "yamada"}
	fmt.Println(i)

	// 要素を削除(keyを指定)
	delete(i, "first_name")
	fmt.Println(i)

	// 値の存在確認(keyを指定)
	v, flg := i["first_name"]
	fmt.Println(flg)
	fmt.Println(v) // string型の空白

	v, flg = i["last_name"]
	fmt.Println(flg)
	fmt.Println(v)
}
