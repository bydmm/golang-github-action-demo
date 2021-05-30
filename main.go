package main

import (
	"fmt"
	"github-action-demo/pets"
)

func main() {
	// 有人会把mysql的密码写代码里并提交到仓库
	var password = "f62e5bcda4fae4f82370da0c6f20697b8f8447ef"
	// 演示代码
	fmt.Printf("my password %s", password)

	saying := pets.Cat()
	fmt.Print(saying)
}
