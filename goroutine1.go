package main

import (
	"fmt"
	"time"
)

func process(n string) {
	time.Sleep(time.Second * 2)
	fmt.Println(n)
}

func main() {

	isFin1 := make(chan bool)
	isFin2 := make(chan bool)
	isFin3 := make(chan bool)

	fmt.Println("Start!")

	go func() {
		process("A")
		isFin1 <- true
	}()
	go func() {
		process("B")
		isFin2 <- true
	}()
	go func() {
		process("C")
		isFin3 <- true
	}()

	// 全部が終わるまでブロックし続ける
	<-isFin1
	<-isFin2
	<-isFin3
	fmt.Println("Finish!")
}
