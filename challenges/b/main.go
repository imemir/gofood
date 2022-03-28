package main

import (
	"fmt"
	"sort"
)

func find(arr []int) int {
	if len(arr)%2 == 0 {
		return 0
	}
	sort.Ints(arr)
	for i := 0; i <= len(arr); i += 2 {
		if len(arr)-1 == i {
			return arr[i]
		} else if arr[i] != arr[i+1] {
			return arr[i]
		}
	}
	return 0
}

func main() {
	arr := []int{2, 2, 5, 6, 5}
	res := find(arr)
	fmt.Println(res)
}
