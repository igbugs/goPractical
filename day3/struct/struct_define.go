package main

import "fmt"

type User struct {
	Name string
	Sex string
	Age int
	AvatarUrl string
}

func NewUser(name string, sex string, age int, avatarusl string) User {
	var user User
	user.Name = name
	user.Sex = sex
	user.Age = age
	user.AvatarUrl = avatarusl

	return user
}

func main() {
	var user User
	user.Name = "user01"
	user.Sex = "male"
	user.Age = 10
	user.AvatarUrl = "http://xxx.com/x.jpg"

	fmt.Printf("user.Name= %s, user.Sex= %s, user.Age= %d\n", user.Name, user.Sex, user.Age)

	user02 := User{
		Name: "user02",
		Age: 18,
		Sex: "male",
	}
	fmt.Printf("user02.Name= %s, user02.Sex= %s, user02.Age= %d\n", user02.Name, user02.Sex, user02.Age)

	user03 := User{}
	fmt.Printf("user03: %#v\n", user03)

	user04 := NewUser("user04","female", 11, "http://xxxxx.cm/xx.bmp")
	fmt.Printf("user04.Name= %s, user04.Sex= %s, user04.Age= %d\n", user04.Name, user04.Sex, user04.Age)

}
