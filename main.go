package main

import (
	"fmt"
)

func Plus(x int, y int) int { //引数と戻り値の型を指定する
	return x + y
}

/*
上と同じ
func Plus(x, y int) int {
	return x + y
}
*/

func Div(x, y int) (int, int) { // Goの特徴の1つ、複数の戻り値を設定できる
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return // 戻り値に変数を指定している為、returnの後を省略出来る
}

func Noreturn() {
	fmt.Println("No return")
	return
}

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a ReturnFunc")
	}
}

func CallFunction(f func()) { //型は func()型
	f() // 引数で受け取った関数を実行
}

/*
	クロージャー(関数閉包)は
	関数が変数として代入された際に、値を保持している関数オブジェクトのことで、
	関数内で定義されている変数などは値は保持する。
  Later()ではstoreの値が保持され、クロージャーとして機能している
*/
func Later() func(string) string {
	var store string // storeの値を空文字で定義
	return func(next string) string {
		s := store // sには空文字
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	//  ---- 名前付関数 ----
	fmt.Println("---- 名前付関数 ----")
	i := Plus(1, 2)
	fmt.Println(i) // => 3

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3) // => 3 0

	i4 := Double(5)
	fmt.Println(i4) // => 10

	Noreturn() // => No return

	//  ---- 無名関数 ----
	fmt.Println("---- 無名関数 ----")
	f := func(x, y int) int {
		return x + y
	}
	i5 := f(1, 2)
	fmt.Println(i5) // => 3

	i6 := func(x, y int) int {
		return x + y
	}(1, 2) // 変数の定義を省略出来る
	fmt.Println(i6)

	//  ---- 関数を返す関数 ----
	fmt.Println("---- 関数を返す関数 ----")
	rf := ReturnFunc()
	rf()                   // => I'm a function
	fmt.Printf("%T\n", rf) // => func()

	//  ---- 関数を引数にとる関数 ----
	fmt.Println("---- 関数を引数にとる関数 ----")
	CallFunction(func() {
		fmt.Println("I'm a CallFunction")
	})

	//  ---- クロージャー(関数閉包) ----
	fmt.Println("---- クロージャー(関数閉包) ----")
	l := Later()
	fmt.Println(l("Hello"))  // => ""
	fmt.Println(l("My"))     // => "Hello"
	fmt.Println(l("name"))   // => "My"
	fmt.Println(l("is"))     // => "name"
	fmt.Println(l("Golang")) // => "is"

	//  ---- ジェネレーター ----
	fmt.Println("---- ジェネレーター ----")

	ints := integers()
	fmt.Println(ints()) // => 1
	fmt.Println(ints()) // => 2
	fmt.Println(ints()) // => 3
	fmt.Println(ints()) // => 4

	otherints := integers()  //代入する変数が変わればクロージャーの領域が変わる
	fmt.Println(otherints()) // => 1
	fmt.Println(otherints()) // => 2
	fmt.Println(otherints()) // => 3
}
