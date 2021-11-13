package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
		3: {Name: "user3", Age: 30},
		4: {Name: "user4", Age: 40},
	}
	fmt.Println(m)
	// => map[1:{user1 10} 2:{user2 20} 3:{user3 30} 4:{user4 40}]

	m2 := map[User]string{
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "LA",
	}
	fmt.Println(m2)
	// => map[{user1 10}:Tokyo {user2 20}:LA]

	m3 := make(map[int]User)
	m3[1] = User{Name: "user3"}
	fmt.Println(m3)
	// => map[1:{user3 0}]

	for _, v := range m {
		fmt.Println(v)
	}
	/*
		=>
			{user1 10}
			{user2 20}
			{user3 30}
			{user4 40}
	*/
}
