package main

import "fmt"

func Cat() string {
	return "Miao~~~~~"
}

func main() {
	saying := Cat()
	fmt.Print(saying)
}
