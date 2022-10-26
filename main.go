package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    `example:"id"`
	Name string `example:"name"`
}

func main() {
	u := User{1, "Andres"}
	GetNumFields(u)
}

func GetNumFields(obj interface{}) {
	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if _, ok := f.Tag.Lookup("example"); ok {
			fmt.Println("Tag found")
			fmt.Println(t.Field(i).Tag.Get("example"))
		} else {
			fmt.Println("Tag not found")
		}
	}
}
