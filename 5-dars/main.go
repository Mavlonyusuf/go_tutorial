package main

import (
	"fmt"
)
 
func main() {
  var mashina = map[string]string{"brand": "KIA", "model": "Carnival", "rang": "Qora", "yili": "2025", "yurgan_yol": ""}
	yangi_mashina := mashina
	fmt.Println(mashina) // map[brand:KIA model:Carnival rang:Qora yili:2025 yurgan_yol:]
	fmt.Println(yangi_mashina) // map[brand:KIA model:Carnival rang:Qora yili:2025 yurgan_yol:]
	delete(yangi_mashina, "yurgan_yol")
	fmt.Println(mashina) // map[brand:KIA model:Carnival rang:Qora yili:2025]
	fmt.Println(yangi_mashina) // map[brand:KIA model:Carnival rang:Qora yili:2025]

}