// arrays
package main
import (
	"fmt"
	"sort"
)

func main()  {
	testArrays1()
	testArrays2()
	testArrays3()
	testArrays4()
	testArrays5()
	testArrays6()
	testSlices1()
	testSlices2()
	testSlices3()
}

func testArrays1() {

	var myarray [3]string

	myarray[0] = "GO"
	myarray[1] = "PROGRAMMING"
	myarray[2] = "ONLINE LEARN"

	// get values of elements using index
	fmt.Println("1", myarray[0])
	fmt.Println('2', myarray[1])
	fmt.Println('3', myarray[2])
}

func testArrays2()  {

	myarray := [3]int{2, 4, 8}
	fmt.Println(myarray)
}

func testArrays3() {

	myarray1 := [...]int{9, 7, 6, 5}
	myarray2 := [4]int{9, 7, 6, 5}

	fmt.Println(myarray1 == myarray2)
}

func testArrays4()  {

	myarray := [3][3]string{{"C#", "GO", "PYTHON"},
	{"Java", "C","C++"},
	{"Dart","swift","Kotlin"}}

	fmt.Println(myarray)
}

func testArrays5()  {
	myarray1 := [3]int{2,4,8}

	myarray2 := myarray1
	fmt.Println(myarray1)
	fmt.Println(myarray2)

	myarray1[2] = 100
	fmt.Println(myarray1)
	fmt.Println(myarray2)
}

func testArrays6()  {
	 myarray1 := [3]int{2,4,8}

	 myarray2 := &myarray1
	 fmt.Println(myarray1)
	 fmt.Println(*myarray2)

	 myarray1[2] = 100
	 fmt.Println(myarray1)
	 fmt.Println(*myarray2)

}

// slices

func testSlices1()  {
	// create new slice
	myslice := []int{2,4,8}

	myslice = append(myslice, 10)

	fmt.Println(len(myslice))
}

func testSlices2()  {
	myarray := [7]string{"1","2","3","4","5","6","7"}

	myslice := myarray[1:4]
	fmt.Println(myslice)

}

func testSlices3()  {
	myslice := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}
	sort.Ints(myslice)
	fmt.Println(myslice)
}

 