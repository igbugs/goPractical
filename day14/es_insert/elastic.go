package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"math/rand"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
	Desc string `json:"desc"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Printf("new client failed, err: %v", err)
		return
	}
	fmt.Println("set client config success")

	for i := 0; i < 10000; i++ {
		p := Person{
			ID:   i,
			Name: fmt.Sprintf("xyb%d", i),
			Age:  rand.Intn(30),
			City: "beijing",
			Desc: "wo wo ha en",
		}

		_, err = client.Index().
			Index("account").
			Type("person").
			BodyJson(p).
			Do(context.Background())
		if err != nil {
			fmt.Printf("insert es failed, err: %v", err)
			return
		}

		fmt.Println("insert es success")
	}
}
