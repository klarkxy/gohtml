package gohtml

import "bytes"

// GoTag: <{name} {property}>{tags}</{name}>
type GoTag struct {
	name        string   // tag名
	property    P        // 属性 xxx=xxx
	tags        []gotext // 里面包含的tag/text
	selfclosing bool     // 是否自闭合
}

func (cur *GoTag) String() string {
	b := new(bytes.Buffer)
	if cur.name != "" {
		b.WriteByte('<')
		b.WriteString(cur.name)
		b.WriteByte(' ')
	}

	for k, v := range cur.property {
		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(v)
		b.WriteByte(' ')
	}
	b.WriteByte('>')
	if !cur.selfclosing {
		// 非自闭合，输出里面的内容和尾巴
		for _, v := range cur.tags {
			b.WriteString(v.String())
			b.WriteByte(' ')
		}

		if cur.name != "" {
			b.WriteString("</")
			b.WriteString(cur.name)
			b.WriteByte('>')
		}
	}

	return b.String()
}

// 在当前Tag中插入若干个Tag/Text，返回值为当前Tag
func (cur *GoTag) Append(txts ...gotext) *GoTag {
	for _, v := range txts {
		cur.tags = append(cur.tags, v)
	}
	return cur
}

// 在当前Tag下新建一个Tag，返回值为新建出来的Tag
func (cur *GoTag) Tag(name string, propertys ...P) *GoTag {
	newtag := NewTag()
	newtag.name = name
	for _, property := range propertys {
		for k, v := range property {
			newtag.property[k] = v
		}
	}
	cur.Append(newtag)
	return newtag
}

// 在当前Tag下新建一个Text，返回值为新建出来的Gotext（其实无意义）
func (cur *GoTag) Text(txt string) *GoText {
	newtext := Text(txt)
	cur.Append(newtext)
	return newtext
}
