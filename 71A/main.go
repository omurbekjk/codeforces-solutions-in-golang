package main

import "fmt"

func Solution(word string) string {
	if len(word) > 10 {
		return fmt.Sprintf("%c%d%c", word[0], len(word)-2, word[len(word)-1])
	} else {
		return word
	}
}

func main() {
	var num int
	fmt.Scanf("%d\n", &num)
	words := make([]string, num)
	answers := make([]string, 0)
	for i := 0; i < num; i++ {
		fmt.Scanf("%s\n", &words[i])
		answers = append(answers, Solution(words[i]))
	}

	for i := 0; i < num; i++ {
		fmt.Printf("%s\n", answers[i])
	}
}