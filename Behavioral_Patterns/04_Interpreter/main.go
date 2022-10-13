package main

func main() {
	expression := "w x z +"
	sentence := NewEvaluator(expression)
	variable := make(map[string]Expression)
	variable["w"] = &Integer{6}
	variable["x"] = &Integer{10}
	variable["z"] = &Integer{41}
	result := sentence.Interpret(variable)
	print(result)
}
