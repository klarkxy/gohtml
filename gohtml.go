/*
 * @Author       : Klarkxy
 * @Date         : 2020-02-07 13:30:31
 * @LastEditors  : Klarkxy
 * @LastEditTime : 2020-02-16 12:32:00
 * @FilePath     : \gohtml\html.go
 */
package gohtml

// 创建一个GoTag对象
func NewTag() *GoTag {
	newtag := GoTag{}
	newtag.attr = make(map[string]string)
	return &newtag
}

func Tag(name string) *GoTag {
	cur := NewTag()
	cur.name = name
	return cur
}
