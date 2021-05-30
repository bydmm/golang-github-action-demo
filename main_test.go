package main

import (
	"github-action-demo/pets"
	"testing"
)

// 🐱是会miao的
func TestCat(t *testing.T) {
	saying := pets.Cat()
	if saying != "Miao~~~~~" {
		t.Errorf("Cat say %s", saying)
	}
}
