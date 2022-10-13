package main

import "fmt"

func main() {
	viewer1 := NewImageViewer("image1.png")
	viewer1.Display()

	viewer2 := NewImageViewer("image1.png")

	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		fmt.Println("not fulfill flyweight")
	}
}
