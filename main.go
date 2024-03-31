package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	fmt.Println(concat("Hello", " world"))
	fmt.Println(concat("Elon", " let's hop this tesla thing workout"))
}
