/*
 * @Author       : Klarkxy
 * @Date         : 2020-02-07 13:30:31
 * @LastEditors  : Klarkxy
 * @LastEditTime : 2020-02-18 16:17:25
 * @FilePath     : \gohtml\gohtml.go
 */
package gohtml

// 创建一个GoTag对象
func NewTag() *GoTag {
	cur := GoTag{}
	cur.attr = make(map[string]string)
	return &cur
}

// 创建一个GoHtml对象
func NewHtml() *GoHtml {
	cur := GoHtml{}
	cur.Init()
	return &cur
}

func Tag(name string) *GoTag {
	cur := NewTag()
	cur.name = name
	return cur
}
