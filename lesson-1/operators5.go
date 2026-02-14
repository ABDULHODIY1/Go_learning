// misc (other) operators

package main 
import "fmt"


func operators5() {
	a := 4

	b := &a
	fmt.Println(*b)
	*b = 7
	fmt.Println(a)

}