package stringendswith

//import "fmt"
import "strings"

func Solution(str, ending string) bool {

	//  fmt.Println("source:",str,ending,"lenght: ", len(str), len(ending))
	//  fmt.Println("'",str[len(str)-len(ending):],"'", " equals ","'",ending,"'" )
	if len(str) >= len(ending) {
		return str[len(str)-len(ending):] == ending
	}
	return false
}

func Solutionbestpractice(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
