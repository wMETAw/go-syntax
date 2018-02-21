package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	fmt.Println("start")

	ch1 := make(chan int)
	ch2 := make(chan string)
	chend := make(chan struct{})

	// channelを送信するgoroutine
	go func(chint chan<- int, chstr chan<- string, end chan<- struct{}) {
		for i := 0; i < 10; i++ {

			// 偶数はint型channel、奇数はstring型channelに送信
			if i%2 == 0 {
				fmt.Println("send to ch1")
				chint <- i
			} else {
				fmt.Println("send to ch2")
				chstr <- "test" + strconv.Itoa(i)
			}
		}

		time.Sleep(time.Second * 1)
		close(end)
	}(ch1, ch2, chend)

	// 受信用ループ
	for {
		select {
		case val := <-ch1:
			fmt.Println("ch1から受信：", val)
		case str := <-ch2:
			fmt.Println("ch2から受信：", str)
		case <-chend:
			fmt.Println("終了")
			return
		}
	}
}
