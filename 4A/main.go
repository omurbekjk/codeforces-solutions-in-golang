package main

import "fmt"

func Solution(w int) string{
	if (w-2)%2 == 0 && (w-2) != 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	w := 0
	fmt.Scanf("%d\n", &w)
	answer := Solution(w)
	fmt.Println(answer)
}
