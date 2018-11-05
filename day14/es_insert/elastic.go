package main

import (
	"github.com/olivere/elastic"
	"fmt"
	"math/rand"
)

type Person struct {
	ID int		`json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
	Desc string `json:"desc"`
}

func main()  {
	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL("http://192.168.20.200"))
	if err != nil {
		fmt.Printf("new client failed, err: %v", err)
		return
	}
	fmt.Printf("set client config success")

	for i := 0; i < 10000; i++ {
		p := Person{
			ID: i,
			Name: fmt.Sprintf("xyb%d", i),
			Age: rand.Intn(30),
			City: "beijing",
			Desc: "wo wo hah en",
		}

		_, err = client.Index().
			Index("account").
			Type("person").
			BodyJson(p).
			Do()
		if err != nil {
			fmt.Printf("insert es failed, err: %v", err)
			return
		}

		fmt.Printf("insert es success")
	}
}