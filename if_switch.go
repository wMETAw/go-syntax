package main

import "fmt"

func main() {

	// 基本的なif
	a := true
	if a {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}

	// ifの中でしか存在しない変数は以下の書き方ができる
	if _score := 60; _score > 80 {
		fmt.Println("Great")
	} else if _score > 60 {
		fmt.Println("Nice")
	} else {
		fmt.Println("Oh...")
	}

	// switch
	signal := "red"
	switch signal {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Caution")
	case "blue":
		fmt.Println("Go")
	default:
		fmt.Println("Accident")
	}

	// switch + if
	score := 80
	switch {
	case score > 80:
		fmt.Println("Great")
	case score > 60:
		fmt.Println("Nice")
	default:
		fmt.Println("Oh...")
	}
}
