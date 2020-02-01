package main

import (
	"fmt"
	"os"

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

	// Example 3
	fmt.Println("Example 3:")
	htm3 := gohtml.NewTag()
	htm3.Tag("table", gohtml.P{"border": "1"}).Func(func(Htm *gohtml.GoTag) {
		Htm.Tag("tr").Tag("td").Text("1").
			Master().Tag("td").Text("2")
		Htm.Tag("tr").Tag("td").Text("3").
			Master().Tag("td").Text("4")
	})
	fmt.Println(htm3.String())

	// Example 4
	fmt.Println("Example 4:")
	htm4 := gohtml.NewTag()
	form := htm4.T("form")
	form.S("input").P("type", "text").P("placeholder", "用户名")
	form.S("input").P("type", "password").P("placeholder", "密码")
	form.T("button").P("type", "submit").Text("登录")
	fmt.Println(htm4.String())

	// Example 5
	fmt.Println("Example 5:")
	htm5 := gohtml.NewTag()
	PWD, _ := os.Getwd()
	htm5.Html("index.html", PWD+"/index.html").Execute("Hello World!")
	fmt.Println(htm5.String())
}
