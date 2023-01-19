package main

import (
	"fmt"
	"sort"
)

// TODO: WA
func C() {
	var n int
	fmt.Scanf("%d", &n)
	odd := []int{}
	even := []int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		if a%2 == 1 {
			odd = append(odd, a)
		} else {
			even = append(even, a)
		}
	}
	sort.Slice(odd, func(i, j int) bool { return odd[i] < odd[j] })
	sort.Slice(even, func(i, j int) bool { return even[i] < even[j] })
	mx := -1
	if len(odd) >= 2 {
		mx = max_value(mx, odd[0]+odd[1])
	}
	if len(even) >= 2 {
		mx = max_value(mx, even[0]+even[1])
	}
	fmt.Println(mx)
}

func max_value(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
