package main

import "fmt"

type Food interface {
	Eat()
}

//Drug -- 抽象的药物接口
type Drug interface {
	Take()
}

//Factory 负责生产药物类和食物类
type Factory interface {
	NewFood() Food
	NewDrug() Drug
}

//-----Food 实现结构体----
type meat struct { // 肉
}

func (t meat) Eat() {
	fmt.Println("Eat meat")
}

type fruit struct { //水果
}

func (t fruit) Eat() {
	fmt.Println("Eat fruit")
}

// -----Drug 实现结构体-----

type feverDrug struct { //发烧药

}

func (t feverDrug) Take() {
	fmt.Println("Take fever drug")
}

type coldDrug struct { //感冒药
}

func (t coldDrug) Take() {
	fmt.Println("Take cold drug")
}

//------ Factory 实现结构体 ------

type FirstFactory struct{}

func (t *FirstFactory) NewFood() Food {
	return meat{}
}

func (t *FirstFactory) NewDrug() Drug {
	return feverDrug{}
}

type SecondFactory struct{}

func (t *SecondFactory) NewFood() Food {
	return fruit{}
}

func (t *SecondFactory) NewDrug() Drug {
	return coldDrug{}
}
