package main

import "fmt"

type Target interface {
	Execute()
}

type Adaptee struct {
}

func (a *Adaptee) SepcificExecute() {
	fmt.Println("充电...")
}

type Adapter struct {
	*Adaptee
}

func (a *Adapter) Execute() {
	a.SepcificExecute()
}
