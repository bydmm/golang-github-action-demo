package main

import "testing"

// 🐱是会miao的
func TestCat(t *testing.T) {
	saying := Cat()
	if saying != "Miao~~~~~" {
		t.Errorf("Cat say %s", saying)
	}
}
