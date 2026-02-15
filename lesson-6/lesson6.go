package main
import (
	"fmt"
	"errors"
	"math"
)


func main()  {
	// res := sum(5,3)
	// fmt.Println(res)

	// res2, err := sqrt(-2)
	// if err != nil{
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res2)
	// }

	var a int = 10
	var b int = 20

	fmt.Println(a,b)
	swap(a, b)
	fmt.Println(a,b)

	fmt.Println(a,b)
	swapByRef(&a, &b)
	fmt.Println(a,b)


}

func sum(x int , y int) int {
	return x + y
}

func sqrt(x float64) (float64, error)  {
	if x < 0 {
		return 0, errors.New("Error!")
	}
	return math.Sqrt(x),nil
}

func swap(x,y int) int  {
	var o int
	o = x
	x = y
	y = o
	return o
}

func swapByRef(x,y *int) int  {
	var o int 
	o = *x
	*x = *y
	*y = o
	return o
}