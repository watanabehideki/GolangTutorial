package main

import (
	"fmt"
	"strconv"
)

func main() {
	//  ---- 整数型(intなど) ----

	// 整数型	     サイズ(bit)	符号	最小値	              最大値
	// int8	  		　8	         あり	 -128	                127
	// int16	  	  16	       あり	 -32768	              32767
	// int32	  	  32         あり	 -2147483648	        2147483647
	// int64	   		64         あり	 -9223372036854775808	9223372036854775807
	// uint8(byte)	8	         なし	 0	                  255
	// uint16	  	　16         なし	 0	                  65535
	// uint32	  	　32         なし	 0	                  4294967295
	// uint64	  	　64         なし	 0	                  18446744073709551615

	var i int = 100 // 環境依存する
	var i2 int64 = 200
	fmt.Println(i + 50)
	// fmt.Println(i + i2) // エラー: invalid operation: i + i2 (mismatched types int and int64)
	fmt.Printf("%T\n", i2)        // %Tは　書式指定子
	fmt.Printf("%T\n", int32(i2)) // => int32 (型変換が可能)

	//  ---- 浮動小数点型(float) ----

	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf) // => +Inf

	ninf := -1.0 / zero
	fmt.Println(ninf) // => -Inf

	nan := zero / zero
	fmt.Println(nan) // => NaN

	//  ---- 論理値型(boolean) ----

	var t, f bool = true, false
	fmt.Println(t, f) // => true false

	//  ---- 文字列型(string) ----

	var s string = "Hello Golang"
	fmt.Println(s) // => Hello Golang
	var si string = "300"
	fmt.Println(si)           // => 300 (string型)
	fmt.Printf("%T\n", si)    // =>  string
	fmt.Println(string(s[0])) // => H (Hello Golang"の1文字目)

	//  ---- byte型(uint8) ----

	byteA := []byte{72, 73}
	fmt.Println(byteA)         // => [72 73]
	fmt.Println(string(byteA)) // => HI (文字列をASCIIコードで表現している)

	c := []byte("HI")
	fmt.Println(c) // => [72 73]

	//  ---- 配列型 ----

	var arr1 [3]int
	fmt.Println(arr1)        // => [0 0 0]
	fmt.Printf("%T\n", arr1) // => [3]int (要素数が型として決まっている)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2) // => [1 2 3]

	arr3 := [...]string{"C", "D"} // [...] は与えれた要素数によって変わる
	fmt.Println(arr3)             // => [C D]
	fmt.Printf("%T\n", arr3)      // => [2]string

	fmt.Println(arr3[0])   // => C
	fmt.Println(arr3[1-1]) // => C

	arr3[0] = "E"
	fmt.Println(arr3) // => [E D]

	fmt.Println(len(arr3)) // => 2 (要素数を返すlenはlengthの意)

	//  ---- interface型 & nil ----

	var x interface{} // intarface型はすべての型と互換性がある
	fmt.Println(x)    // => <nil>

	// x = 1
	// fmt.Println(x)     // => 1
	// fmt.Println(x + 2) // エラー: nvalid operation: x + 2 (mismatched types interface {} and int)

	//  ---- 型変換 ----

	var fl_change_type float64 = 10.7
	i_change_type := int(fl_change_type)
	fmt.Println(i_change_type)                        // => 10
	fmt.Printf("i_change_type = %T\n", i_change_type) // => i_change_type = int

	var s2 string = "100"
	fmt.Printf("s2 = %T\n", s2) // => st = string

	i3, _ := strconv.Atoi(s2)
	fmt.Println(i3)             // => 100
	fmt.Printf("i3 = %T\n", i3) // => i3 = int

	var i4 int = 200
	s3 := strconv.Itoa(i4)
	fmt.Println(s3)
	fmt.Printf("s3 = %T\n", s3) // => s3 = string

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b) // => [72 101 108 108 111 32 87 111 114 108 100]

	h2 := string(b)
	fmt.Println(h2) // => Hello World
}
