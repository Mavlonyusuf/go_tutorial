package main

import "fmt"

func main() {
	myPhone := phone{Make: "Apple", Model: "17 pro max", Color: "natural titanium", Weight: 221, Price: 2300, Processor: processor{Chip: "A18 Pro", RAM: 12}}

	fmt.Println(myPhone)
}
type phone struct {
	Make, Model, Color string
	Weight, Price      int
	Processor processor
}
type processor struct {
	Chip string
	RAM int
}