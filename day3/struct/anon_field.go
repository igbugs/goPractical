package main

import "fmt"

type User struct {
	Name string
	Sex string
	Age int
	AvatarUrl string
	int			// 匿名字段
	string		// 匿名字段
}

func main() {
	var user User
	user.Name = "user01"
	user.Sex = "male"
	user.Age = 10
	user.AvatarUrl = "http://xxx.com/x.jpg"
	user.int = 110
	user.string = "Hello"

	fmt.Printf("user.Name= %s, user.Sex= %s, user.Age= %d\n", user.Name, user.Sex, user.Age)
	fmt.Println(user.int, user.string)

}
