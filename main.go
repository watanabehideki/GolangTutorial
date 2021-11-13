package main

import (
	"fmt"
)

func Double(i int) {
	i = i * 2
}

func Double2(i *int) { // 参照渡し
	*i = *i * 2
}

func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n) // メモリ上のアドレス
	// => 0xc0000b2008

	Double(n)
	fmt.Println(n)
	// => 100

	var p *int = &n
	// *int でポインタ型をm指定
	// &n でアドレスを渡す

	fmt.Println(p)
	// => 0xc0000b2008
	fmt.Println(*p)
	// => 100

	// *p = 300
	// // *p と n は同じアドレスなので上書きされる
	// fmt.Println(n)
	// // => 300

	// n = 200
	// fmt.Println(*p)
	// // => 200

	Double2(&n)
	fmt.Println(n)
	// => 200

}
