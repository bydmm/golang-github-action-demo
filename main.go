package main

import "fmt"

func Cat() string {
	return "Wang~~~~~"
	// return "Miao~~~~~"
}

func main() {
	saying := Cat()
	fmt.Print(saying)
}
