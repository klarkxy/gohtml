/*
 * @Author       : Klarkxy
 * @Date         : 2020-02-07 13:30:31
 * @LastEditors  : Klarkxy
 * @LastEditTime : 2020-02-16 12:29:51
 * @FilePath     : \gohtml\tag.go
 */
package gohtml

import (
	"bytes"
	"fmt"
)

// Tag: <{name} {attr}}>{tags}</{name}>
type GoTag struct {
	name       string            // tag名
	attr       map[string]string // 属性列表
	singleAttr []string          // 无值属性
	tags       []fmt.Stringer    // 所有支持String()输出的都可以作为tag
	self       bool              // 自闭合
}

func (cur *GoTag) String() string {
	b := new(bytes.Buffer)
	if cur.name != "" {
		b.WriteByte('<')
		b.WriteString(cur.name)

		for k, v := range cur.attr {
			b.WriteByte(' ')
			b.WriteString(k)
			b.WriteString("=\"")
			b.WriteString(v)
			b.WriteString("\"")
		}
		for _, k := range cur.singleAttr {
			b.WriteByte(' ')
			b.WriteString(k)
		}
		b.WriteByte('>')
	}
	if !cur.self {
		// 非自闭合，输出里面的内容和尾巴
		for _, v := range cur.tags {
			b.WriteString(v.String())
		}
		if cur.name != "" {
			b.WriteString("</")
			b.WriteString(cur.name)
			b.WriteByte('>')
		}
	}
	return b.String()
}

// 在当前Tag中插入若干个Tag/Text，返回当前Tag
func (cur *GoTag) Append(tags ...fmt.Stringer) *GoTag {
	for _, v := range tags {
		cur.tags = append(cur.tags, v)
	}
	return cur
}

// 在当前Tag下新建一个Tag，返回新建出来的Tag
func (cur *GoTag) Tag(name string) *GoTag {
	newtag := Tag(name)
	cur.Append(newtag)
	return newtag
}

// 设置当前Tag为self-Closing，返回当前Tag
func (cur *GoTag) Self() *GoTag {
	cur.self = true
	return cur
}

// 在当前Tag中插入一条string，返回当前Tag
func (cur *GoTag) Text(str string) *GoTag {
	buf := bytes.NewBufferString(str)
	cur.Append(buf)
	return cur
}

// 设置属性，返回当前Tag
func (cur *GoTag) Attr(attrs ...string) *GoTag {
	for i := 0; i+1 < len(attrs); i += 2 {
		cur.attr[attrs[i]] = attrs[i+1]
	}

	return cur
}

// 设置属性，返回当前Tag
func (cur *GoTag) SingleAttr(attrs ...string) *GoTag {
	for i := 0; i+1 < len(attrs); i++ {
		cur.singleAttr = append(cur.singleAttr, attrs[i])
	}

	return cur
}
