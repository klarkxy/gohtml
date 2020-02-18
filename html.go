/*
 * @Author       : Klarkxy
 * @Date         : 2020-02-18 15:26:54
 * @LastEditors  : Klarkxy
 * @LastEditTime : 2020-02-18 18:38:24
 * @FilePath     : \gohtml\html.go
 */
package gohtml

import (
	"strings"
)

type GoHtml struct {
	html   *GoTag
	meta   *GoTag
	link   *GoTag
	head   *GoTag
	body   *GoTag
	script *GoTag
}

// 初始化
func (cur *GoHtml) Init() *GoHtml {
	cur.html = Html()
	head := cur.html.Head()
	body := cur.html.Body()
	cur.meta = head.Tag("")
	cur.link = head.Tag("")
	cur.head = head.Tag("")
	cur.body = body.Tag("")
	cur.script = body.Tag("")
	return cur
}

func (cur *GoHtml) String() string {
	str := "<!DOCTYPE html>\n" + cur.html.String()
	for {
		str2 := str
		str2 = strings.ReplaceAll(str2, "\n\n", "\n")
		str2 = strings.ReplaceAll(str2, " \n", "\n")
		str2 = strings.ReplaceAll(str2, "> ", ">")
		if str2 == str {
			break
		}
		str = str2
	}
	return str
}
func (cur *GoHtml) MinString() string {
	str := cur.String()
	for {
		str2 := str
		str2 = strings.ReplaceAll(str2, "> ", ">")
		str2 = strings.ReplaceAll(str2, " <", "<")
		str2 = strings.ReplaceAll(str2, ">\n", ">")
		str2 = strings.ReplaceAll(str2, "\n<", "<")
		if str2 == str {
			break
		}
		str = str2
	}

	return str
}

// 获取Html页签，用来设置一些全局属性
func (cur *GoHtml) Html() *GoTag {
	return cur.html
}

// 获取Head页签
func (cur *GoHtml) Head() *GoTag {
	return cur.head
}

// 获取Body页签
func (cur *GoHtml) Body() *GoTag {
	return cur.body
}

// 新增一条<meta>
func (cur *GoHtml) Meta() *GoTag {
	return cur.meta.Meta()
}

// 新增一条<link>
func (cur *GoHtml) Link() *GoTag {
	return cur.link.Link()
}

// 新增一条<script></script>
func (cur *GoHtml) Script() *GoTag {
	return cur.script.Script()
}
