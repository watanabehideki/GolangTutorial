package main

import "fmt"

const Pi = 3.14 //頭文字が大文字だと他のパッケージからも飛び出しが可能

const (
	URL      = "http://example.com"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "あ"
	E
	F
)

const (
	c0 = iota //連番を返す
	c1
	c2
)

func main() {
	fmt.Println(Pi)            // => 3.14
	fmt.Println(URL, SiteName) // => http://example.com test

	fmt.Println(A, B, C, D, E, F) // => 1 1 1 あ あ あ

	fmt.Println(c1)         // => 1
	fmt.Println(c0, c1, c2) // => 0 1 2

}
