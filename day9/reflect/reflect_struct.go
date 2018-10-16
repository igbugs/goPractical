package main

import (
	"reflect"
	"fmt"
	)

type User struct {
	Name string	`json:"name"`
	Age int		`json:"age"`
	Sex string	`json:"sex"`
}

func testStructValue(in interface{})  {
	v := reflect.ValueOf(in)
	t := v.Type()

	switch t.Kind() {
	case reflect.Struct:
		fieldNum := t.NumField()
		fmt.Printf("field number: %d\n", fieldNum)
		for i := 0; i < fieldNum; i++ {
			field := t.Field(i)
			vField := v.Field(i)

			fmt.Printf("field[%d] name: %s, json key: %s, val: %v\n",
				i, field.Name, field.Tag.Get("json"), vField.Interface())
		}
	}
}

func main() {
	var user = User{
		Name: "xyb",
		Age: 10,
		Sex: "F",
	}

	testStructValue(user)
	fmt.Printf("user: %#v\n", user)
}
