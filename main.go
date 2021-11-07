package main

import (
	"fmt"
	"github.com/kravetsd/go-katas/nbyear"
	"github.com/kravetsd/go-katas/reversestrings"
	"github.com/kravetsd/go-katas/stringendswith"
)

func main() {
	fmt.Println(nbyear.NbYear(1500, 5, 100, 5000))
	fmt.Println(reversestrings.Solution("test"))
	fmt.Println(stringendswith.Solution("test", "st"))
}
