package main

import "fmt"

func Cat() string {
	word := "Miao"
	return fmt.Sprintf("%s~~~~~", word)
}

func main() {
	saying := Cat()
	fmt.Print(saying)
}
