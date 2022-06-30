package main

import (
	"fmt"
)

func main() {

	a := binsearch.search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5)
	fmt.Println(a)
}
