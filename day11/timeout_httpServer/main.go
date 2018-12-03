package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		num := rand.Intn(2)
		if num == 0 {
			time.Sleep(time.Second * 5)
			fmt.Fprintf(w, "slow request")
			return
		}
		fmt.Fprintf(w, "quick request")
	})

	http.ListenAndServe(":8080", nil)
}
