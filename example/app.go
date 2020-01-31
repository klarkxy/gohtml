package main

import (
	"fmt"

	"gitee.com/Klarkxy/gohtml"
)

func main() {
	// Example 1
	fmt.Println("Example 1:")
	htm1 := gohtml.NewTag()
	htm1.Tag("p").Text("Hello World!")
	fmt.Println(htm1.String())

	// Example 2
	fmt.Println("Example 2:")
	htm2 := gohtml.NewTag()
	htm2.Tag("a", gohtml.P{"href": "www.baidu.com"}).Text("前往百度")
	fmt.Println(htm2.String())
}
