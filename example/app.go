package main

import (
	"fmt"

	"gitee.com/Klarkxy/gohtml"
)

func main() {
	html := gohtml.NewTag()

	fmt.Println(html.String())
}
