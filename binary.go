package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target := []int{2, 8, 15, 12, 25}
	res := make([]int, len(target))
	var wg sync.WaitGroup
	for i := 0; i < len(target); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res[i] = DevideNDomination(arr, 0, len(arr)-1, target[i])
		}(i)
	}
	wg.Wait()
	fmt.Println(res)
}
func DevideNDomination(arr []int, lpoint, rpoint, target int) int {
	if (rpoint-lpoint) < 2 && arr[rpoint] != target && arr[lpoint] != target {
		return -1
	}
	if arr[rpoint] == target {
		return rpoint
	}
	if arr[lpoint] == target {
		return lpoint
	}
	if arr[(rpoint+lpoint)/2] > target {
		rpoint = (rpoint + lpoint) / 2
	} else {
		lpoint = (rpoint + lpoint) / 2
	}
	return DevideNDomination(arr, lpoint, rpoint, target)
}
