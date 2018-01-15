package main

import (
	"fmt"
	"time"
)

func loading(ret chan string) {

	fmt.Printf("Now Loading")

	// channelに値を送信
	time.Sleep(time.Millisecond * 800)
	ret <- "."

	time.Sleep(time.Millisecond * 400)
	ret <- "."

	time.Sleep(time.Millisecond * 500)
	ret <- "."

	close(ret)
}

func main() {

	// channelのインスタンス生成
	ret := make(chan string)
	go loading(ret)

	for i := range ret {
		fmt.Printf(i)
	}
}
