// go basic operators
package main
import "fmt"

func operators2() {

	number := 39
	number2 := 90

	if number == number2{
		fmt.Println("True")

	}else if  number != number2 {
		fmt.Println("false")

	}else if !(number == number2) {
		fmt.Println("True")
	}

}