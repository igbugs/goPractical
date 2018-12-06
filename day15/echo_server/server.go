package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	templ *template.Template
	addr = flag.String("addr", "localhost:8080", "http service address")

	// use default options
	upgrader = websocket.Upgrader{}
)

func main()  {
	flag.Parse()
	log.SetFlags(0)

	err := initTempl()
	if err != nil {
		log.Println("init templ failed, err: ", err)
		return
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/echo", echo)
	log.Printf("start websocket server in %s ...\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func initTempl() (err error) {
	templ, err = template.ParseFiles("./views/index.html")
	if err != nil {
		log.Fatalf("parse index.html failed, err:%v", err)
	}
	return
}

func home(w http.ResponseWriter, r *http.Request)  {
	_ = templ.Execute(w, "ws://" + r.Host + "/echo")
}

func echo(w http.ResponseWriter, r *http.Request)  {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println("close connection failed, err:", err)
		}
	}()

	for {
		//mt, msg, err := conn.ReadMessage()
		//if err != nil {
		//	log.Println("read:", err)
		//	break
		//}
		//
		//log.Println("recv: ", msg)
		//err = conn.WriteMessage(mt, msg)
		//if err != nil {
		//	log.Println("write: ", err)
		//	break
		//}

		err = conn.WriteMessage(websocket.TextMessage, []byte("this is server push message"))
		if err != nil {
			log.Println("push msg failed, err: ",err)
		}
		time.Sleep(time.Second)
	}


}