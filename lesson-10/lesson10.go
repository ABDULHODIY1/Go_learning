// WaitGroup

package main

import (
	"fmt"
	"time"
	"sync"
)

// func main()  {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go display("Hello",&wg)
// 	go display("Hi", &wg)
	
// 	wg.Wait()
// }

// func display(input string, wg *sync.WaitGroup) {
// 	for i :=1; i <= 7; i++ {
// 		fmt.Println(i, input)
// 		time.Sleep(1 * time.Second)
// 	}
// 	wg.Done()
// }


func main()  {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		display("Hello")
		wg.Done()}()
	go func() {
		display("Hi")
		wg.Done()}()
	
	wg.Wait()
	fmt.Println("Complected")
}

func display(input string) {
	for i :=1; i <= 7; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}