package main

import (
	"fmt"
)

func chanfunc() {
	//チャネルとは、複数のゴルーチン間でデータの送受信をするためのデータ構造
	// チャネルはキュー(先入れ先出し)の性質を持っている

	//チャネルの宣言
	var ch1 chan int

	//受信専用
	// var ch2 <-chan int

	//送信専用
	// var ch3 chan<- int

	//チャネルの初期化(writing, readingを可能にする)
	ch1 = make(chan int)

	ch2 := make(chan int)

	fmt.Println("# チャネルの容量(バッファサイズ)")
	fmt.Println(cap(ch1)) // => 0
	fmt.Println(cap(ch2)) // => 0

	ch3 := make(chan int, 5) //バッファサイズの指定（バッファサイズ分の要素数しか入れられない）
	fmt.Println(cap(ch3))    // => 5

	fmt.Println("# チャネルのデータを送信")
	ch3 <- 2
	fmt.Println(len(ch3))
	// => 1 (lenでデータ数を表示)
	ch3 <- 3
	ch3 <- 4
	fmt.Println("要素数", len(ch3)) // => 3

	fmt.Println("# チャネルからデータを受信")
	i := <-ch3
	fmt.Println(i) // => 2(156行目で最初に入れた値)
	fmt.Println("要素数", len(ch3))
	i2 := <-ch3
	fmt.Println(i2) // => 3(159行目で2番目に入れた値)
	fmt.Println("要素数", len(ch3))

	fmt.Println(<-ch3)
	fmt.Println("要素数", len(ch3))
}

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}
func closefunc() {
	ch1 := make(chan int, 2)
	close(ch1)

	// ch1 <- 1

	i, ok := <-ch1 //第２引数はチャネル内のバッファが空かどうかの真偽値が入る。(closeすると受信はできないが、送信はできる。)
	fmt.Println(i, ok)
	// => 0 false
}

func reciever2(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func chanfor() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1) // チャネルでfor文を使用するときは、closeが必須
	for i := range ch1 {
		fmt.Println(i)
	}

}

func main() {
	fmt.Println("---- チャネル ----")
	chanfunc()
	fmt.Println("# チャネルとゴルーチン")
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// go reciever(ch1)
	// go reciever(ch2)

	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	ch2 <- i
	// 	time.Sleep(50 * time.Millisecond)
	// 	i++
	// }
	fmt.Println("# チャネルのclose")
	// ch1 := make(chan int, 2)
	// closefunc()

	// go reciever2("1.goroutin", ch1)
	// go reciever2("2.goroutin", ch1)
	// go reciever2("3.goroutin", ch1)

	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	i++
	// }
	// close(ch1)
	// time.Sleep(3 * time.Second)
	fmt.Println("# チャネルのfor文")
	chanfor()
	fmt.Println("# チャネルのselect文")
	// チャネルはある1つのチャネルがデッドロックになると、それ以降に記述されチャネルには到達しなくなる。

	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch2 <- "あ"
	// select文のcaseにはチャネルの処理を記述する必要がある。また、処理されるcaseはランダム。
	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

}
