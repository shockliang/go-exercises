package main

import (
	"fmt"
	"generics-exercises/priorityqueue"
)

func main() {
	queue := priorityqueue.NewPriorityQueue[int, string]([]int{priorityqueue.High, priorityqueue.Medium, priorityqueue.Low})

	queue.Add(priorityqueue.Low, "L-1")
	queue.Add(priorityqueue.High, "H-1")

	fmt.Println(queue.Next())

	queue.Add(priorityqueue.Medium, "M-1")
	queue.Add(priorityqueue.High, "H-2")
	queue.Add(priorityqueue.High, "H-3")

	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
}
