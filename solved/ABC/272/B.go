package main

import "fmt"

func B() {
	var n, m int

	fmt.Scanf("%d %d", &n, &m)

	// TODO: ２次元配列をFalseで初期化するプログラムあり
	xf := make([][]bool, n)

	for i := 0; i < n; i++ {
		for l := 0; l < n; l++ {
			xf[i] = append(xf[i], false)
		}
	}

	for i := 0; i < m; i++ {
		var k int
		fmt.Scanf("%d", &k)

		x := make([]int, k)

		for j := 0; j < k; j++ {
			fmt.Scanf("%d", &x[j])
			x[j] = x[j] - 1
		}

		for ii := 0; ii < len(x)-1; ii++ {
			for jj := ii + 1; jj < len(x); jj++ {
				xf[x[ii]][x[jj]] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !xf[i][j] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")

}
