package main

func main() {
	strategyA := NewStrategyA()
	c := NewContext()
	c.SetStrategy(strategyA)
	c.Execute()
	strategyB := NewStrategyB()
	c.SetStrategy(strategyB)
	c.Execute()
}
