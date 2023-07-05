package main

import "fmt"

func main() {
	var sentence = "The quick brown fox jumped over the alphabet soup"
	sentence = leftTab(sentence, 5, " ")
	fmt.Println(sentence)
	sentence = delimit(sentence, " ")
	fmt.Println(sentence)

}
