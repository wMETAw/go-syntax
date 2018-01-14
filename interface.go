package main

import "fmt"

// インターフェイス
type greeter interface {
	greet()
}

// JPN構造体とメソッド
type jpn struct{}

func (j jpn) greet() {
	fmt.Println("こんにちは")
}

type usa struct{}

func (u usa) greet() {
	fmt.Println("Hello")
}

func main() {

	// greeter型スライスを作成
	greeters := []greeter{jpn{}, usa{}}

	for _, v := range greeters {
		v.greet()
		checkInterface(v)
	}
}

/*
 * 型チェック
 * t interface{} は空のインターフェースであり、全ての型を受け取ることができる
 */
func checkInterface(t interface{}) {

	/**
	 * 型アサーション
	 * 値, フラグ := t.(type) // t が type を満たすかを調べる
	 */
	_, flg := t.(jpn)
	if flg {
		fmt.Println("I am")
	} else {
		fmt.Println("I'm")
	}

	// 型switch
	switch t.(type) {
	case jpn:
		fmt.Println("Japanese")
	case usa:
		fmt.Println("American")
	default:
		fmt.Println("Human")
	}
}
