package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

type Sample2 struct {
	ID   int     `json:"id"`
	Body *Sample `json:"body,omitempty"`
}

func main() {
	fmt.Println("run")
	arg := "test"
	optionalFunc(&arg)
	optionalFunc(nil)
	arg2 := Sample{}
	optionalStruct(&arg2)
	optionalStruct(nil)
	arg3 := Sample2{}
	optionalStruct2(&arg3)
	optionalStruct2(nil)
	getTags(&arg3)
}

func optionalFunc(arg *string) error {
	if arg == nil {
		fmt.Println("optionalFunc: nil")
		return nil
	}
	fmt.Printf("optionalFunc: %s\n", *arg)
	return nil
}

func optionalStruct(arg *Sample) error {
	if arg == nil {
		fmt.Println("optionalStruct: nil")
		return nil
	}
	jsonstr, _ := json.Marshal(arg)
	fmt.Printf("optionalStruct: %s\n", string(jsonstr))
	return nil
}

func optionalStruct2(arg *Sample2) error {
	if arg == nil {
		fmt.Println("optionalStruct2: nil")
		return nil
	}
	jsonstr, _ := json.Marshal(arg)
	fmt.Printf("optionalStruct2: %s\n", string(jsonstr))
	return nil
}

func getTags(arg *Sample2) error {
	if arg == nil {
		fmt.Println("getTags: nil")
		return nil
	}
	t := reflect.TypeOf(*arg)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s\n", field.Tag.Get("json"))
	}
	return nil
}
