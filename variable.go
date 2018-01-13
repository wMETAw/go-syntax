package main

import(
	"fmt"
)

func main(){
	a := 5          // int
	b := 1.35       // float
	c := "hoge"     // string
	var d bool      // bool
	const e = "定数" // 定数

	fmt.Printf("a:%d b:%f c:%s d:%t e:%s \n", a,b,c,d,e)

	// 定数列挙
	// 識別子iota(イオタ)を使用することで0から始まる連番にできる
	const (
		sun = iota + 1
		mon
		tue
	)
	fmt.Println(sun, mon, tue)
}