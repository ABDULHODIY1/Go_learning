// Select statement

package main
import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "Fast"
			time.Sleep(time.Millisecond * 600)

		}
	}()

	go func() {
		for {
			channel2 <- "Slow"
			time.Sleep(time.Second * 2)

		}

	}()
	for {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)

		case message2 := <-channel2:
			fmt.Println(message2)
		
		// default:
		// 	fmt.Println("Error: No info")
		}
	
	}

}