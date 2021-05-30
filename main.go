package main

import (
	"fmt"
	"github-action-demo/pets"
)

var (
	// 有人会把mysql的密码写代码里并提交到仓库
	password = "123456"
)

func main() {
	// 演示代码
	fmt.Printf("my password %s", password)

	saying := pets.Cat()
	fmt.Print(saying)
}
