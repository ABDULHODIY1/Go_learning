// theme

package main
import "fmt"

func Filter[T any](items []T, predicate func(T)bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}


type Users struct {
	Name string
	Age int
}

func main(){
	evens := Filter([]int{1,2,3,4,5}, IsEven)
	fmt.Println(evens)

	adults := Filter([]Users{
		{"Alice",16},
		{"Tony",19},
		{"Jhon",17},
	}, func (u Users) bool  {
		return u.Age >= 18
	})

	fmt.Println(adults)
}
func IsEven(n int) bool {
	return n%2 == 0
}