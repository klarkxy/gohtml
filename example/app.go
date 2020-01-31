package main

import (
	"fmt"

	"gitee.com/Klarkxy/gohtml"
)

func main() {
	// Example 1
	fmt.Println("Example 1:")
	htm1 := gohtml.Tag("p").Text("Hello World!")
	fmt.Println(htm1.String())

	// Example 2
	fmt.Println("Example 2:")
}
