package httpServer

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello world %s!</h1>", r.FormValue("name"))
	})

	http.ListenAndServe(":8080", nil)

}
