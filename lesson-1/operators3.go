// bitwise operators
package main
import "fmt"

func operators3()  {
	a := 14
	b := 15


	// | OR
	res1 := a | b
	fmt.Println(res1)

	// & AND
	res2 := a & b
	fmt.Println(res2)

	// ^ XOR
	res3 := a ^ b
	fmt.Println(res3)

	// << left shift
	res4 := a << b
	fmt.Println(res4)

	// >> Right Shift
	res5 := a >> b
	fmt.Println(res5)

	// &^ AND NOT
	res6 := a &^ b
	fmt.Println(res6)
	
}