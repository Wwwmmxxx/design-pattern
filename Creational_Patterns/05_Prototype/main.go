package main

import "fmt"

var manager *PrototypeManager

func init() {
	manager = NewPrototypeManager()

	t1 := &Type1{
		name: "type1",
	}
	manager.Set("t1", t1)
}

func main() {
	Clone()

	CloneFromManager()
}

func Clone() {
	t1 := manager.Get("t1")

	t2 := t1.Clone()

	if t1 == t2 {
		fmt.Println("error! get clone not working")
	}
}

func CloneFromManager() {
	c := manager.Get("t1").Clone()

	t1 := c.(*Type1)
	if t1.name != "type1" {
		fmt.Println("error")
	}

}
