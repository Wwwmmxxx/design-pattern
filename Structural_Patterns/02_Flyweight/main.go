package main

import "fmt"

func main() {
	factory := NewFlyWeightFactory()
	hong := factory.GetFlyWeight("hong beauty")
	xiang := factory.GetFlyWeight("xiang beauty")
	fmt.Println(hong, xiang)
}
