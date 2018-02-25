package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func main() {
	result := testing.Benchmark(func(b *testing.B) {
		run("A", "B", "C", "D", "E")
	})
	fmt.Println(result.T)
}

func run(name ...string) {
	fmt.Println("start!")

	// waigroupの作成
	wg := new(sync.WaitGroup)

	// channnelの数だけ作成
	isFin := make(chan bool, len(name))

	for _, v := range name {

		// 1処理に対して１つ増やす
		wg.Add(1)

		// サブスレッドに処理を任せる
		go process(v, isFin, wg)
	}
	wg.Wait()
	close(isFin)
	fmt.Println("Finish!")
}

func process(name string, isFin chan bool, wg *sync.WaitGroup) {

	// wgをデクリメント
	defer wg.Done()
	time.Sleep(time.Second * 2)
	fmt.Println(name)
	isFin <- true
}
