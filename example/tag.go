/*
 * @Author       : Klarkxy
 * @Date         : 2020-02-07 13:30:31
 * @LastEditors  : Klarkxy
 * @LastEditTime : 2020-02-16 12:41:44
 * @FilePath     : \gohtml\example\app.go
 */
package main

import (
	"fmt"

	"github.com/klarkxy/gohtml"
)

func main() {
	{
		// Example 1
		fmt.Println("Example 1:")
		htm := gohtml.Tag("p").Text("Hello World!")
		fmt.Println(htm.String())
	}
	{
		// Example 2
		fmt.Println("Example 2:")
		htm := gohtml.Tag("a").Attr("href", "https://www.zhihu.com/").Text("打开知乎")
		fmt.Println(htm.String())
	}
	{
		// Example 3
		fmt.Println("Example 3:")
		htm1 := gohtml.Tag("button").Attr("type", "button").Attr("href", "https://www.zhihu.com/").Text("打开知乎")
		htm2 := gohtml.Tag("button").Attr("type", "button", "href", "https://www.zhihu.com/").Text("打开知乎")
		fmt.Println(htm1.String())
		fmt.Println(htm2.String())
	}
	{
		// Example 4
		fmt.Println("Example 4:")
		htm1 := gohtml.Tag("meta").Attr("http-equiv", "X-UA-Compatible").Attr("content", "IE=edge").Self()
		fmt.Println(htm1.String())
	}
	{
		// Example 5
		fmt.Println("Example 5:")
		table := gohtml.Tag("table")
		tr1 := table.Tag("tr")
		tr1.Tag("th").Text("1")
		tr1.Tag("th").Text("2")
		tr2 := table.Tag("tr")
		tr2.Tag("th").Text("3")
		tr2.Tag("th").Text("4")
		fmt.Println(table.String())
	}
	{
		// Example 6
		fmt.Println("Example 6:")
		form := gohtml.Form()
		form.H2().Text("Please Sign In")
		form.Input().Type("username").Placeholder("Username").Required().Autofocus()
		form.Input().Type("password").Placeholder("Password").Required()
		divlabel := form.Div().Label()
		divlabel.Input().Type("checkbox").Value("Remember me")
		divlabel.Text("Remember me.")
		form.Button().Type("submit").Text("Sign in")
		fmt.Println(form.String())
	}
}
