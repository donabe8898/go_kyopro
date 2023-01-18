package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 0; i < n; i++ {
		ans += a[i]
	}
	fmt.Println(ans)
}
