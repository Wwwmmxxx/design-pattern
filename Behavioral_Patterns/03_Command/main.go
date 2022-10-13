package main

func main() {
	c1 := ConcreteCommand1{&Receiver1{}}
	c2 := ConcreteCommand1{&Receiver2{}}
	c3 := ConcreteCommand1{&Receiver3{}}

	i := Invoker{}
	i.SetCommand(&c1)
	i.SetCommand(&c2)
	i.SetCommand(&c3)
	i.ExecuteCommand()
}
