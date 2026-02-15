// map in go
// map[keyType]ValTY

package main
import "fmt"

func main()  {

	status := make(map[string]int)

	// add value to map
	status["pending"] = 0
	status["inited"] = 1
	status["running"] = 2
	status["timeout"] = 3
	status["successful"] = 4
	status["failed"] = 5

	// read values from map
	var successfulStatus = status["successful"]
	fmt.Println(successfulStatus)

	// delete the value from the map
	delete(status, "timeout")
}