package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- sum(n[i])
		}(i)
	}
	res1 := <-ch
	res2 := <-ch
	res3 := <-ch
	fmt.Println("result:", res1+res2+res3)
	wg.Wait()
}

func sum(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}
