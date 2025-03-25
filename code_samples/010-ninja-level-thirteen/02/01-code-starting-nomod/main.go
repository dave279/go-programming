package main

import (
	"./quote"
	"./word"
	"fmt"
	//	"ninja13-02start.quote"
	//	"ninja13-02start.word"
	//
	// "github.com/dave279/go-programming/code_samples/010-ninja-level-thirteen/02/01-code-starting/quote"
	// "github.com/dave279/go-programming/code_samples/010-ninja-level-thirteen/02/01-code-starting/word"
)

func main() {
	//	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
