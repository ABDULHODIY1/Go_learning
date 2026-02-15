// struct looka like class
package main
import "fmt"

type car struct {
	Make, Color, Model string
	Year, Weight int
	Engine engine
	
}
type engine struct{
	name string
	hp int
}

func main()  {
	//var MyCar Car
	myCar :=car{Make: "BMW", Model:"M8", Color: "Black", Year: 2024, Weight: 3021, Engine: engine{"V8", 1200}}
	fmt.Println(myCar)
	// fmt.Println(myCar.Model)
	
}