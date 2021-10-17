package main

import (
	"fmt"
)

/*
配列とスライスの違い
- 配列は要素数が決まっているが、スライスは要素数が可変で拡張可能
append() -拡張する関数
*/

func appendfunc() {
	sl := []int{100, 200}
	fmt.Println(sl)
	// => [100 200]

	sl = append(sl, 300) //300を最後に追加
	fmt.Println(sl)
	// => [100 200 300]
	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)
	// => [100 200 300 400 500 600]
}
func makefunc() {
	sl := make([]int, 5)
	fmt.Println(sl)
	// => [0 0 0 0 0]
	fmt.Println(len(sl))
	// => 5
}

func capfunc() {
	// cap()は変数のメモリ使用量を返す関数
	sl := make([]int, 5, 10) //make(型, 要素数, 容量)
	fmt.Println(cap(sl))
	// => 10
	sl = append(sl, 1, 2, 3, 4, 5, 6)
	fmt.Println(cap(sl))
	// => 20 (要領以上の要素が追加されるとメモリの消費が倍になってしまう。)
}

func copyfunc() {
	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	n := copy(sl2, sl) // copy(コピー先, コピー元) copy()はコピーした要素数を返す
	fmt.Println(n, sl2)
	// => 5 [1 2 3 4 5]
}

func slfor() {
	sl := []string{"a", "b", "c"}
	fmt.Println(sl)

	for i := range sl {
		fmt.Println(i)
	}

	for i2 := 0; i2 < len(sl); i2++ {
		fmt.Println(sl[i2])
	}
}

func Sum(s ...int) int {
	n := 0                //初期値を設定
	for _, v := range s { //第１引数のインデックスはパーシャル
		n += v // sに入る値が1つずつ取り出されて、vに代入される
	}
	return n
}
func mapfunc() {
	var m = map[string]int{"A": 100, "B": 200}

	fmt.Println(m)
	// => map[A:100 B:200]

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)
	// => map[A:100 B:200]

	m3 := make(map[int]string)
	fmt.Println(m3)
	// => []

	m3[1] = "JAPAN"
	m3[2] = "USA"
	fmt.Println(m3)
	// => map[1:JAPAN 2:USA]

	fmt.Println(m["A"])
	// => 100
	fmt.Println(m["C"])
	// => 0 (int型の初期値0が返る。string型だと、空文字が返る。)

	fmt.Println("# エラーハンドリング")
	s, ok := m3[3]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s, ok)
	// => 空文字 false

	fmt.Println("# 値の更新")
	m3[2] = "CANADA"
	fmt.Println(m3)
	// => map[1:JAPAN 2:CANADA]

	m3[3] = "CHINA"
	fmt.Println(m3)
	// => map[1:JAPAN 2:CANADA 3:CHINA]

	fmt.Println("# 値の削除")
	delete(m3, 3)
	fmt.Println(m3)
	// => map[1:JAPAN 2:CANADA]

}

func mapforfunc() {
	m := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}
	for k, v := range m {
		fmt.Println(k, v)
		// => Apple 100
		// 		Banana 200
	}
}

func main() {
	var sl []int
	fmt.Println(sl)
	// => []
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)
	// => [100, 200]

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)
	// => [A B]

	// ---- make関数 ----
	sl4 := make([]int, 5)
	fmt.Println(sl4)
	// => [0 0 0 0 0]

	// ---- 値の代入 ----
	sl2[0] = 1000 // 変数名[インデックス番号] = 変更する値
	fmt.Println(sl2)
	// => [1000 200]

	// ---- 値の取り出し ----
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)
	fmt.Println(sl5[3])
	// => 4

	// ---- 値の取り出し(範囲指定) ----
	fmt.Println(sl5[2:4]) //インデックス番号2〜４まで
	// => [3 4]
	fmt.Println(sl5[:2]) //インデックス番号2まで
	// => [1 2]
	fmt.Println(sl5[2:]) //インデックス番号2以降から最後まで
	// => [3 4 5]
	fmt.Println(sl5[:]) //全て
	// => [1 2 3 4 5]
	fmt.Println(sl5[len(sl5)-1]) //len()はjsのlengh()と同じで、配列内の要素数を返す。ここでは5を返し、5-1のインデックス番号が4番目の値を返す。
	// => 5
	fmt.Println(sl5[1 : len(sl5)-1]) //インデックスの最初と最後を取り除く
	// => [2 3 4]

	fmt.Println("---- append関数 ----")
	appendfunc()
	fmt.Println("---- make関数 ----")
	makefunc()
	fmt.Println("---- cap関数 ----")
	capfunc()
	fmt.Println("---- copy関数 ----")
	copyfunc()
	fmt.Println("---- スライスのfor文 ----")
	slfor()
	fmt.Println("---- 可変長引数 ----")
	fmt.Println(Sum(1, 2, 3))
	// => 6
	sl6 := []int{1, 2, 3, 4}
	fmt.Println(Sum(sl6...))
	// => 10 (スライスも渡せる)
	fmt.Println("---- マップ ----")
	mapfunc()
	fmt.Println("---- マップのfor文 ----")
	mapforfunc()
}
