package main

import (
	"fmt"
	"time"
)

func task1() {

	// 二秒間停止
	time.Sleep(time.Second * 2)
	fmt.Println("task1 finish")
}

func task2() {
	fmt.Println("task2 finish")
}

func main() {

	// ゴルーチン（並行処理開始）
	go task1()
	go task2()

	// task1が終了する前にmainが終わるため３秒まつ
	time.Sleep(time.Second * 3)
}
