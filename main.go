package main

import (
	"fmt"
)

// 構造体(struct)とはオブジェクト指向のクラスのようなもの
// 複数の型の値を1つにまとめたもの

// User型の構造体の定義
type User struct {
	Name string
	Age  int
	// X, Y int (同じ型のフィールドはまとめて定義できる)
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	// => { 0} (各フィールドの初期値が出力される)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)
	// => {user1 10}

	// 暗黙的な変数定義
	user2 := User{}
	fmt.Println(user2)
	// => { 0}
	user2.Name = "user2"
	user2.Age = 20
	fmt.Println(user2)
	// => {user2 20}

	// 初期値を持たせて定義
	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)
	// => {user3 30}

	// フィールドを指定しないで初期値を設定
	user4 := User{"user4", 40}
	fmt.Println(user4)
	// => {user4 40}

	// Nameだけ指定
	user5 := User{Name: "user5"} //フィールド名は必須
	fmt.Println(user5)
	// => {user5 0}

	user6 := new(User) // new(User)は構造体のポインタ型を返す
	fmt.Println(user6)
	// => &{ 0} (ポインタ型なのでアドレス演算子"&"がつく)

	user7 := &User{}
	fmt.Println(user7)
	// => &{ 0} (new(User) と同じ結果)

	UpdateUser(user1)
	UpdateUser2(user7)

	fmt.Println(user1)
	// => {user1 10} (UpdateUserは値型を引数に受け取っているので、更新されない)
	fmt.Println(user7)
	// => &{A 1000} (UpdateUser2はポインタ型を引数に受け取っているので、アドレスが参照されるので更新される)

}
