// channels

package main
import (
	"fmt"
	// "math/rand"
	"time"
)


// func main()  {
// 	channel := make(chan string, 2)
// 	channel <- "Apple"
// 	channel <- "banana"

// 	inf := <- channel 
// 	inf1 := <- channel 
// 	fmt.Println(inf, inf1)
// }

// func main()  {
// 	channel := make(chan string)

// 	go func(){
// 		channel <- "Apple"
// 		channel <- "banana"
// 	}()
// 	inf := <- channel 
// 	inf1 := <- channel 
// 	fmt.Println(inf, inf1)
// }

// func main() {
// 	channel := make(chan int)
// 	go GetRandomNumber(channel)
// 	for randomNumber := range channel {
// 		fmt.Println("random Numbers:", randomNumber)
// 	}
// }

// func GetRandomNumber(channel chan int)  {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 1 ; i <= 3; i++ {
// 		number := rand.Intn(1000)
// 		time.Sleep(time.Second * 1)
// 		channel <- number
// 	}
// 	close(channel)
// }

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "Fast"
			time.Sleep(time.Millisecond * 100)

		}
	}()

	go func() {
		for {
			channel2 <- "Slow"
			time.Sleep(time.Second * 2)

		}

	}()
	for {
		fmt.Println(<-channel1)
		fmt.Println(<-channel2)
	}

}