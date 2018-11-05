package main

import (
	"github.com/satori/go.uuid"
	"fmt"
)

func main()  {
	uid := uuid.Must(uuid.NewV4())
	fmt.Println("UUID: ", uid.String())

}
