package main

func main() {
	first := &FirstFactory{} //这里用的&符号,因为只实现了指针类型
	first.NewFood().Eat()
	first.NewDrug().Take()

	second := &SecondFactory{}
	second.NewFood().Eat()
	second.NewDrug().Take()
}
