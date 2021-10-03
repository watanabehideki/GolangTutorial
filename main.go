package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func anything(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!!")
	case int:
		fmt.Println(v + 2)
	}
}

func TestDefer() {
	defer fmt.Println("END") // fmt.Println("START")の後に実行される
	fmt.Println("START")
}
func RunDefer() {
	defer fmt.Println("あ")
	defer fmt.Println("い")
	defer fmt.Println("う")
}

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func init() { // main()の前に実行される
	fmt.Println("init関数")
}

func main() {
	//  ---- if文 ----
	fmt.Println("---- if文 ----")

	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	fmt.Println("---- 簡易文付if文 ----")
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; x == 2 {
		fmt.Println("xは2です")
	}
	fmt.Println(x) // => 0

	//  ---- エラーハンドリング ----
	fmt.Println("---- エラーハンドリング ----")

	var s string = "文字列"
	i, err := strconv.Atoi(s) // 引数で受け取ったstrirg型の数字をint型として返す
	if err != nil {
		fmt.Println(err) // => strconv.Atoi: parsing "文字列": invalid syntax
	}
	fmt.Printf("i = %T\n", i)

	//  ---- for文 ----
	fmt.Println("---- for文 ----")

	i2 := 0
	for {
		i2++
		if i2 == 3 {
			break
		}
		fmt.Printf("i2は%dです\n", i2)
		/*
			=>
				i2は1です
				i2は2です
		*/
	}
	count := 0
	for count < 4 {
		fmt.Println(count)
		count++
		/*
			=>
				0
				1
				2
				3
		*/

	}
	//上と同じ
	for count2 := 0; count2 < 4; count2++ {
		fmt.Println(count2)
	}
	// ある条件と一致する時、処理をスキップする
	for count3 := 0; count3 < 4; count3++ {
		if count3 == 2 {
			continue
		}
		fmt.Println(count3)
		/*
			=>
				0
				1
				3
		*/
	}
	fmt.Println("---- 配列のfor文 ----")

	arr := [3]string{"い", "ろ", "は"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		/*
			=>
				い
				ろ
				は
		*/
	}
	for i, v := range arr {
		fmt.Println(i, v)
		/*
			=>
				0 い
				1 ろ
				2 は
		*/
	}
	fmt.Println("---- ラベル付のfor ----")

Loop:
	for {
		for {
			for {
				fmt.Println("START")
				break Loop
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("処理をしない")
	}
	fmt.Println("END")

Loop2:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop2
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}

	//  ---- switch文 ----
	fmt.Println("---- 式のswitch ----")

	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}
	// 上と同じ
	switch n := 1; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	switch { // caseに演算子を使う場合はswitchの後は書かない(書くとエラーになる)
	case n > 0:
		fmt.Println("n > 0")
	case n < 0:
		fmt.Println("n < 0")
	default:
		fmt.Println("I don't know")
	}

	fmt.Println("---- 型のswitch ----")

	var x2 interface{} = "xは"
	switch x2.(type) {
	case int:
		fmt.Println("int型です")
	case string:
		fmt.Println("string型です")
	default:
		fmt.Println("I don't know")
	}
	switch v := x2.(type) {
	case int:
		fmt.Println(v, "int型です")
	case string:
		fmt.Println(v, "string型です")
	default:
		fmt.Println(v, "I don't know")
	}

	anything(1)   // => 3
	anything("あ") // => あ!!

	//  ---- defer ----
	/*
		- 関数が終了した時の処理を設定できる
		- 1つの関数内で複数設定されている時は、登録される逆順に実行される

	*/
	fmt.Println("---- defer ----")

	TestDefer()

	/*
		defer func() { // RunDefer()の後に実行される
			fmt.Println("1")
			fmt.Println("2")
			fmt.Println("3")
		}()
	*/

	RunDefer()

	file, err := os.Create("defer.txt") // => defer.txtを作成
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("Hello defer")) // => defer.txtにHello deferを記述

	//  ---- panic & recover ----
	fmt.Println("---- panic & recover ----")
	/*
		defer func() { // panicで強制的に例外を発生させdeferで例外処理を実行
			if x := recover(); x != nil {
				fmt.Printf("x = %s\n", x) // => x = runtime error
			}
		}()
		panic("runtime error") // panic以降は実行されない
		fmt.Println("実行されない")  // 実行されない
	*/

	// ---- ゴルーチン(並行処理) ----
	fmt.Println("---- ゴルーチン(並行処理) ----")
	/*
		go sub() // goをつけることによって、以下のfor文と同時に実行される
		for {
			fmt.Println("Main Loop")
			time.Sleep(200 * time.Millisecond)
		}
	*/

	// ---- init ----
	/*
		- main()より先に実行される
		- (ほぼ使わないが)複数設定できる

	*/
	fmt.Println("---- init ----")

	fmt.Println("Main")

}
