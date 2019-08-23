package main

import "fmt"

type fetch map[string]string

func main() {
	v := fetch{
		"test":  "test",
		"test1": "test",
		"test2": "test",
		"test3": "test",
	}

	for key, value := range v {
		fmt.Println(key, value)
	}
}
