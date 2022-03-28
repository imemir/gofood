package main

import "fmt"

func generate(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func squaring(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) int {
	var out int
	for n := range in {
		out += n
	}
	return out
}

func main() {
	nums := []int{12, 54, 89, 21, 66, 47, 14, 285, 96}
	sum := sum(squaring(generate(nums)))
	fmt.Println(sum)
}
