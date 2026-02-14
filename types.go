// basic types

package main
import "fmt"

func types(){
	// Integer

	var x uint8 = 225
	fmt.Println(x +1 , x)

	var y int16 = 32765
	fmt.Println(y - 2)
	
	//float

	a := 1.3
	b := 34.34
	c := b - a

	fmt.Println("Result :", c)
	fmt.Println("result of c : %T", c) // %T it's show a type of the value


	// complex

	var d complex128 = complex(6,2)
	var w complex64 = complex(9,2)

	fmt.Println(d,"\n", w)

	// Boolean

	str1 := "Learn-Types"
	str2 := "learn-Types"
	str3 := "Learn-types"

	result1 := str1 == str2
	result2 := str1 == str3
	result3 := str1 == str1

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println("type of result-1 %T and type of result-2 and result-3 type.", result1, result2, result3)

	//string

	str := "Learn-Types"
	fmt.Println("\nstr text leght %d", len(str))
	fmt.Println("\nText: %s", str)
	fmt.Println("\n type of str: %T", str)

	// constantly variables

	const PI = 3.14

	fmt.Println(PI)
}