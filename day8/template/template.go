package main

import (
	"html/template"
	"fmt"
	"net/http"
)

var t *template.Template

type Address struct {
	City string
	Province string
	PostCode string
}

type User struct {
	Name string
	Age int
	Address
}

func initTemlate() (err error) {
	t, err = template.ParseFiles("C:/GoProject/Go3Project/src/day8/template/index.html")
	if err != nil {
		fmt.Printf("load temlate file is failed, err: %v\n", err)
		return
	}
	return
}

func handleUserInfo(w http.ResponseWriter, r *http.Request) {
	var users []*User
	for i := 0; i < 10; i++ {
		var user = User{
			Name: "user01",
			Age: 12,
			Address: Address{
				City: "beijing",
				Province: "beijing",
				PostCode: "110101",
			},
		}

		users = append(users, &user)
	}

	t.Execute(w, users)
	_ = r
}

func main() {
	err := initTemlate()
	if err != nil {
		return
	}

	http.HandleFunc("/user/info", handleUserInfo)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Listen failed, err: %v\n", err)
		return
	}
}
