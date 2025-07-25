package main

import "fmt"

func main() {
	// fmt.Println(oddOrEven(11))
	// fmt.Println(returnCube(8))
	// fmt.Println(posOrNeg(-465))
	// fmt.Println(bigOrLit(12, 25))
	// var myPhone phone
	myPhone := phone{Make: "Apple", Model: "17 pro max", Color: "natural titanium", Weight: 221, Price: 2300}

	fmt.Println(myPhone)
}

type phone struct {
	Make, Model, Color string
	Weight, Price      int
}

// 1
// func oddOrEven(son1 int) string {
// 	if son1 % 2 == 0 {
// 		return "Bu son juft"
// 	} else {
// 		return "Bu son toq"
// 	}
// }
// // 2
// func returnCube (son int) int {
// 	return son * son * son
// }
// // 3
// func posOrNeg(son int) string {
// 	if son > 0  {
// 		return "Ushbu son musbat"
// 	}else {
// 		return "Ushbu son manfiy"
// 	}
// }
// 4
// func bigOrLit(son1 int, son2 int) string {
// 	if son1 > son2 {
// 		return "Birinchi son katta"
// 	}else if son1 == son2 {
// 		return "Ikki son ham teng"
// 	}else {
// 		return "Ikkinchi son katta"
// 	}
// }