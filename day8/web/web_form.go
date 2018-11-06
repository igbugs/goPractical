package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func userLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("C:\\GoProject\\Go3Project\\src\\day8\\web\\edit.html")
		//data, err := ioutil.ReadFile("./edit.html")
		if err != nil {
			http.Redirect(w, r, "./404.html", http.StatusNotFound)
			return
		}

		w.Write(data)
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "admin" {
			fmt.Fprintf(w, "Login success")
		} else {
			fmt.Fprintf(w, "Login failed")
		}
	}
}

func main() {
	http.HandleFunc("/user/login", userLogin)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
