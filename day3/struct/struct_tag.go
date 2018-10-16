package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	Name string	`json:"name"`
	Sex string	`json:"sex"`
	Age int		`json:"age"`
	AvatarUrl string	`json:"avatar_url"`
}

func main() {
	var user User
	user.Name = "user01"
	user.Sex = "male"
	user.Age = 10
	user.AvatarUrl = "http://xxx.com/x.jpg"

	fmt.Printf("user.Name= %s, user.Sex= %s, user.Age= %d\n", user.Name, user.Sex, user.Age)

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("marshal failed, err %v\n", err)
		return
	}
	fmt.Printf("json: %v\n", string(data))
}
