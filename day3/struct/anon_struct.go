package main

import "fmt"

type Address struct {
	City string
	Province string
}

type User struct {
	Name string
	Sex string
	Age int
	AvatarUrl string
	City string
	int			// 匿名字段
	string		// 匿名字段
	Address		// 匿名结构体
}

func main() {
	var user User
	user.Name = "user01"
	user.Sex = "male"
	user.Age = 10
	user.AvatarUrl = "http://xxx.com/x.jpg"
	user.int = 110
	user.string = "Hello"

	user.City = "HeNan"
	user.Address.City = "BeiJing"
	user.Address.Province = "BeiJing"

	user02 := User{
		int: 1000,
		string: "world",
		Address: Address{
			City: "beijing",
			Province: "shanghai",
		},
	}


	fmt.Printf("user.Name= %s, user.Sex= %s, user.Age= %d\n", user.Name, user.Sex, user.Age)
	fmt.Println(user.int, user.string, user.Address, user.City, user.Address.City)
	fmt.Println(user02.int, user02.string, user02.Address)
}
