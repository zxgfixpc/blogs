package main

import (
	"fmt"

	"blogs/httpserver"
)

// 项目启动入口
func main() {
	//
	fmt.Println("starting...")
	httpserver.Start()
}
