package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City   string
	Street string
}

type Student struct { // tag: `json:"rename"`  omitempty:省略为空的字段
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"age,omitempty"`
	Address Address `json:"address,omitempty"`
}

func main() {
	s := Student{
		Name: "liu bo",
		Age:  18,
		Address: Address{
			City:   "Beijing",
			Street: "Changping",
		},
	}

	// transfer struct to json
	bytes, e := json.Marshal(s)

	if e != nil {
		panic(e)
	}

	fmt.Printf("%s\n", bytes)

	// transfer json to struct
	var s1 Student
	e = json.Unmarshal([]byte(bytes), &s1)
	if e != nil {
		panic(e)
	}
	fmt.Printf("%v\n", s1)
}
