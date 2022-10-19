package main

import "fmt"

func main() {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do()

	if res != "pre:real:after" {
		fmt.Println("error occur")
	}
}
