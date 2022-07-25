package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	for i := 0; i < 3; i++ {
		go func(i int) {
			ch <- sum(n[i])
		}(i)
	}
	res1 := <-ch
	res2 := <-ch
	res3 := <-ch
	fmt.Println("result:", res1+res2+res3)
}

func sum(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}
