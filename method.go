package main

import "fmt"

type user struct {
	first_name string
	last_name  string
	score      int
}

func main() {

	u := user{first_name: "Taro", last_name: "Yamada", score: 89}
	u.increment()
	u.show()
	fmt.Println(u)
}

// メソッド
func (u *user) increment() {
	u.score++
}

func (u *user) show() {
	fmt.Printf("Name:%s%s Score:%d", u.first_name, u.last_name, u.score)
}
