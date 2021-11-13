package main

import (
	"fmt"
)

type T struct {
	User
}

type User struct {
	Name string
	Age  int
}

// structメソッド
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u *User) SetName(name string) { //構造体に定義したフィールドの値を変更する際はレシーバーをポインタ型にする必要がある
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()
	// => user1

	//Nameを更新する
	user1.SetName("A")
	user1.SayName()
	// => A

	// 構造体の埋め込み
	t := T{User: User{Name: "userT", Age: 10}}
	fmt.Println(t)
	// => {{userT 10}}

	fmt.Println(t.User)
	// => {userT 10}

	fmt.Println(t.User.Name)
	// => userT

	fmt.Println(t.Name) // t.User.Name
	// => userT (type T でUser型を省略記述した場合のみ、t.User.Nameと同じ結果が得られる)

	t.User.SetName("t")
	fmt.Println(t.User)
	// => {t 10}
}
