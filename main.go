package main

import (
	"fmt"
	"github-action-demo/pets"
)

func main() {
	// 有人会把mysql的密码写代码里并提交到仓库
	// 所以应该从环境变量中取
	var password = "dsfdstfewr345324534532rdrsft4354w5"
	// 演示代码
	fmt.Printf("my password %s", password)

	saying := pets.Cat()
	fmt.Print(saying)
}
