package main

import (
	"fmt"

	"github.com/kravetsd/go-katas/binsearch"
)

func main() {

	a := binsearch.Search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5)
	fmt.Println(a)
}
