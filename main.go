package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var s []int

	var e int
	for i := 0; i < x; i++ {
		fmt.Scan(&e)
		s = append(s, e)
	}

	var suma int
	for _, v := range s {
		suma = suma + v
	}
	fmt.Println(suma)
}
