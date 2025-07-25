package main

import (
	"fmt"
)

func main (){

	_, b := add(5, "Hello")
	fmt.Println(b) 
	fullInfo("Mavlon", "Yusuf", 24)
// arifmetik operatorlar
// num1 :=  60
// num2 := 15

// fmt.Println(num1 + num2) // 75
// fmt.Println(num1 - num2) // 45
// fmt.Println(num1 * num2) // 900
// fmt.Println(num1 / num2) // 4
// fmt.Println(num1 % num2) // 0
// num1++
// fmt.Println("num1: " , num1) // num1:  61
// num2--
// fmt.Println("num2: " , num2) // num2:  14

// 



}
func fullInfo(firstName string, lastName string, age uint16) {
	fmt.Printf("Salom mening ismim %s %s, yoshim %d da.", firstName, lastName, age)
}
func add(x int, y string) (natija int, hisobot string) {
  natija = x + x
  hisobot = y + " World!"
  return
}