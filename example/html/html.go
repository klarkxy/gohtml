/*
 * @Author       : Klarkxy
 * @Date         : 2020-02-18 16:02:06
 * @LastEditors  : Klarkxy
 * @LastEditTime : 2020-02-18 18:40:54
 * @FilePath     : \gohtml\example\html.go
 */
package main

import (
	"fmt"

	"github.com/klarkxy/gohtml/bootstrap"

	"github.com/klarkxy/gohtml"
)

func main() {
	{
		htm := gohtml.NewHtml()
		fmt.Println("String:")
		fmt.Println(htm.String())
		fmt.Println("MinString:")
		fmt.Println(htm.MinString())
	}
	fmt.Println("")
	{
		bootstrap := gohtml.NewHtml()
		bootstrap.Html().Lang("zh-CN")
		// Meta部分
		bootstrap.Meta().Charset("utf-8")
		bootstrap.Meta().Http_equiv("X-UA-Compatible").Content("IE=edge")
		bootstrap.Meta().Name("viewport").Content("width=device-width, initial-scale=1")
		// Css引入
		bootstrap.Link().Href("https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css").Rel("stylesheet")
		// Js引入
		bootstrap.Script().Src("https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js")
		bootstrap.Script().Src("https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js")
		// Head
		bootstrap.Head().Title().Text("Bootstrap 101 Template")
		// Body
		bootstrap.Body().H1().Text("Hello World!")

		fmt.Println(bootstrap.String())
	}
	fmt.Println("")
	{
		bootstrap := bootstrap.Bootstrap()
		// Head
		bootstrap.Head().Title().Text("Bootstrap4 101 Template")
		// Body
		bootstrap.Body().H1().Text("Hello World!")
		fmt.Println(bootstrap.String())
	}
}
