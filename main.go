package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type Users []*User

//コンストラクタ関数
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	// structのコンストラクタ
	user1 := NewUser("user1", 10)
	fmt.Println(user1)
	// => &{user1 10}

	fmt.Println(*user1) //user1をデリファレンス
	// => {user1 10}

	// structのスライス
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}
	user5 := User{Name: "user5", Age: 50}

	users := Users{}
	users = append(users, &user2)
	users = append(users, &user3)
	users = append(users, &user4, &user5)

	fmt.Println(users)
	// => [0xc00000c078 0xc00000c090 0xc00000c0a8 0xc00000c0c0]
	fmt.Println("users =>")
	for _, u := range users {
		fmt.Println(*u)
	}
	/*
		=>
				{user2 20}
				{user3 30}
				{user4 40}
				{user5 50}
	*/

	users2 := make([]*User, 0)
	users2 = append(users2, &user4, &user5)
	fmt.Println("users2 => ")
	for _, u := range users2 {
		fmt.Println(*u)
	}
	/*
		=>
				{user4 40}
				{user5 50}
	*/

}
