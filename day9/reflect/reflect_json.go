package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string
}

func marshal(data interface{}) (jsonStr string) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	switch t.Kind() {
	case reflect.Int, reflect.String, reflect.Int32:
		jsonStr = fmt.Sprintf("\"%v\"", data)
	case reflect.Struct:
		numField := t.NumField()
		for i := 0; i < numField; i++ {
			name := t.Field(i).Name
			tag := t.Field(i).Tag.Get("json")
			if len(tag) > 0 {
				name = tag
			}

			vField := v.Field(i)
			vFieldValue := vField.Interface()

			if t.Field(i).Type.Kind() == reflect.String {
				jsonStr += fmt.Sprintf("\"%s\":\"%s\"", name, vFieldValue)
			} else {
				jsonStr += fmt.Sprintf("\"%s\":%v", name, vFieldValue)
			}

			if i != numField-1 {
				jsonStr += ","
			}
		}
		jsonStr = "{" + jsonStr + "}"
	}
	return
}

func main() {
	var user = User{
		Name: "xyb",
		Age:  12,
		Sex:  "M",
	}

	jsonStr := marshal(user)
	fmt.Printf("user marshal output: %s", jsonStr)
}
