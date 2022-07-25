package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	var resoult int
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	for i := 0; i < len(n); i++ {
		go func(i int) {
			ch <- sum(n[i])
		}(i)
	}

	for i := 0; i < len(n); i++ {
		resoult += <-ch
	}
	fmt.Println("result:", resoult)
}

func sum(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}
