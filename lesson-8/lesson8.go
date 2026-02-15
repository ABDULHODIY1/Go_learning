package main

import "fmt"

func main()  {
	TestQueue()
	
}

func TestQueue() {
	queue := NewQueue()
	queue.Enqueue(100)

	fmt.Println("Queue", queue)
	queue.Enqueue(25).Enqueue(77)

	fmt.Println("Queue", queue)
	fmt.Println("is queue empty?", queue.IsEmpty())

	var resault, _ = queue.Peek()
	fmt.Println("The next word item in the queue:", resault)

	fmt.Println("Dequeue one item....")
	resault, _ = queue.DeQueque()
	
	fmt.Println(resault)
	fmt.Println("Dequeue one item....")
	resault, _ = queue.DeQueque()
	
	fmt.Println(resault)
}

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Peek() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.data[0], nil
}

func (q *Queue) Enqueue(n int) *Queue {
	q.data = append(q.data, n)
	return q
}

func (q *Queue) DeQueque() (int, error)  {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queque is Empty")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}