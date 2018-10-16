package main

import "fmt"

type Address struct {
	City  string
	Province  string
	CreateTime string
}

type Email struct {
	Account  string
	CreateTime string
}

type User struct {
	Username  string
	Sex  string
	Age  int
	AvatarUrl string
	Address
	Email
	CreateTime string
}

func main() {
	user := User{
		Username: "liming",
		Sex: "female",
		Age: 19,
		//CreateTime: "20180712",
		Address: Address{
			City: "beijing",
			CreateTime: "20180713",
		},
		Email: Email{
			Account: "xueyongbo@163.com",
			CreateTime: "20180713",
		},
	}

	fmt.Println(user.Address.CreateTime)
}
