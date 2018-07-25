package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("Frank", "Frank Larios")
	fmt.Println(match)
	match, _ = regexp.MatchString("Larios", "Frank larios")
	fmt.Println(match)

	parse, err := regexp.Compile("[Ff]rank")
	if err != nil {
		fmt.Printf("Error compiling RE: %s\n", err)
	} else {
		fmt.Println(parse.MatchString("Frank Larios"))
		fmt.Println(parse.MatchString("frank Larios"))
		fmt.Println(parse.MatchString("F rank Larios"))
		fmt.Println(parse.ReplaceAllString("frank Frank", "FRANK"))
	}
}
