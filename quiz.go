package main

import "fmt"

func main() {
	// Problem is Banana add to original slice also
	/*fruits := []string{"Apple", "Orange", "Plum", "Pomegranate", "Grape"}
	some := fruits[1:3]
	some = append(some, "Banana")

	fmt.Println("Fruits :", fruits)
	fmt.Println("some : ", some)*/

	fruits := []string{"Apple", "Orange", "Plum", "Pomegranate", "Grape"}
	// Solution1
	/*some := make([]string, 2)
	copy(some, fruits[1:3])*/
	some := fruits[1:3:3]

	some = append(some, "Banana")
	fmt.Println("Fruits :", fruits)
	fmt.Println("some : ", some)

}
