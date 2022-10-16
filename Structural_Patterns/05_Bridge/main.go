package main

func main() {
	// Common - SMS
	m1 := NewCommonMessage(ViaSMS())
	m1.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via SMS

	m2 := NewCommonMessage(ViaEmail())
	m2.SendMessage("have a drink?", "bob")

	m3 := NewUrgencyMessage(ViaSMS())
	m3.SendMessage("have a drink?", "bob")

	m4 := NewUrgencyMessage(ViaEmail())
	m4.SendMessage("have a drink?", "bob")
}
