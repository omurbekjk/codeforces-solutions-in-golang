package main

import "fmt"

func Solution(n, m, a int64) int64 {
	var height = (n + a - 1) / a
	var width = (m + a - 1) / a
	return height * width
}

func main() {
	var n, m, a int64
	fmt.Scanf("%d %d %d\n", &n, &m, &a)
	answer := Solution(n, m, a)
	fmt.Println(answer)
}
