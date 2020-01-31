package main

import (
	"fmt"

	"gitee.com/Klarkxy/gohtml"
)

func main() {
	html := gohtml.NewTag()
	html.Tag("p").Text("Hello World!")
	fmt.Println(html.String())
}
