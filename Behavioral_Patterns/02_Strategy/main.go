package main

func main() {
	// Pay by Cash
	payment1 := NewPayment("Ada", "", 123, &Cash{})
	payment1.Pay()

	// Pay by Bank
	payment2 := NewPayment("Bob", "0002", 888, &Bank{})
	payment2.Pay()
}
