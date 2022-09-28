package main

func main() {
	MeatFactory{}.NewFood(MeatKind).Eat()   //去肉厂吃肉
	FruitFactory{}.NewFood(FruitKind).Eat() //去水果厂吃水果
	VegetableFactory{}.NewFood(VegetableKind).Eat()
	NutFactory{}.NewFood(NutKind).Eat()
}
