package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K, X int
	var num int
	fmt.Scan(&N, &K, &X)

	//長さに合わせて配列を用意し格納
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	for i := 0; i < N; i++ {

		num = A[i] / X

		if K >= num {
			A[i] -= (num * X)
			K -= num
		} else {
			A[i] -= (K * X)
			K = 0
			break
		}

	}

	//割引できる金額が大きい順にソート
	sort.Sort(sort.Reverse(sort.IntSlice(A)))

	//残りのクーポンを、ソートした順で使う
	for i := 0; K > 0; i++ {
		K -= 1
		A[i] = 0

		if i == N-1 {
			break
		}
	}

	sum := 0

	for _, x := range A {
		sum += x
	}

	fmt.Printf("%d", sum)

}
