package main

import "fmt"

func outer() {
	outer := "outer関数です!"
	fmt.Println(outer)
}

func main() {
	// # fmtの使い方
	// fmt.Println("Hello World")
	// fmt.Println(time.Now())

	// # 変数と型
	var i int = 100
	fmt.Println(i)
	i = 150
	fmt.Println(i) // => 150

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3) // => 0 （初期値として 0, 空文字が入る為)

	// # 暗黙的な定義
	i4 := 400 // var i4 int = 400 と同じ
	fmt.Println(i4)

	// i4 = "Hello"
	// fmt.Println(i4) // => エラー: cannot use "Hello" (type untyped string) as type int in assignment

	outer() // => "outer関数です!"

	// s5 := "Not Use" //エラー: s5 declared but not used
}
