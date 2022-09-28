package main

func main() {
	worker := NewWorker(&Coder{})
	worker.Daily()
}
