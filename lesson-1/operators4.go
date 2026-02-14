package main
import "fmt"

func operators4() {	
	var a int = 35
	var b int = 40

	a = b
	fmt.Println(a)

	a += b
	fmt.Println(a)

	a -= b
	fmt.Println(a)

	a /= b 
	fmt.Println(a)
	
	a %= b
	fmt.Println(a)

	a *= b
	fmt.Println(a)

}