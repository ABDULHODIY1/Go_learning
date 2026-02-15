// for loop

package main
import (
	"fmt"
	"time"
	"strconv"
)

func main()  {
	for i:=1; i<8;i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("infinity loop" + strconv.Itoa(time.Now().Second()))
		time.Sleep(1 * time.Second)
		break
	}

	i := 1
	for i < 10 {
		fmt.Println(i)
		i++

	}

	// myarr := [3]string{"Yellow","Red","Blue"}
	// for index, value := range myarr {
	// 	fmt.Println("index",index,"value",value)
	// }

	mymap := map[int]string {
		1:"Go",
		2:"Python",
		3:"JavaScript",
	}
	for key, value := range mymap {
		fmt.Println("index",key,"value",value)
	}
}