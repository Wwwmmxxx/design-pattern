package main

import "fmt"

type IReceiver interface {
	Action()
}

type Receiver1 struct {
}

func (r *Receiver1) Action() {
	fmt.Println("执行命令1")
}

type Receiver2 struct {
}

func (r *Receiver2) Action() {
	fmt.Println("执行命令2")
}

type Receiver3 struct {
}

func (r *Receiver3) Action() {
	fmt.Println("执行命令3")
}

type ICommand interface {
	Execute()
}

type ConcreteCommand1 struct {
	receiver IReceiver
}

func (c *ConcreteCommand1) Execute() {
	c.receiver.Action()
}

type ConcreteCommand2 struct {
	receiver IReceiver
}

func (c *ConcreteCommand2) Execute() {
	c.receiver.Action()
}

type ConcreteCommand3 struct {
	receiver IReceiver
}

func (c *ConcreteCommand3) Execute() {
	c.receiver.Action()
}

type Invoker struct {
	commands []ICommand
}

func (i *Invoker) SetCommand(command ICommand) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommand() {
	for _, c := range i.commands {
		c.Execute()
	}
}
