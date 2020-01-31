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

		for k, v := range cur.property {
			b.WriteString(k)
			b.WriteString("=\"")
			b.WriteString(v)
			b.WriteString("\" ")
		}
		b.WriteByte('>')
	}
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
	b.WriteByte('\n')

	return b.String()
}

// 在当前Tag中插入若干个Tag/Text，返回值为当前Tag
// 谨慎调用，可能会出现重复的情况
func (cur *GoTag) Append(txts ...gotext) *GoTag {
	for _, v := range txts {
		cur.tags = append(cur.tags, v)
	}
	return cur
}

// 与Append的区别在于，Inc实际上不会进行任何操作，在参数计算的时候会自动完成插入功能
// 推荐用法：html.Inc(html.Tag("xxx"), html.Text("yyy"))
func (cur *GoTag) Inc(txts ...GoText) *GoTag {
	return cur
}

// 更复杂的嵌套，用函数来实现吧
// 该函数没有返回值，通常作为一条语句的最末端，需要返回值用FuncR
func (cur *GoTag) Func(fn func(*GoTag)) {
	fn(cur)
}

// 更复杂的嵌套，用函数来实现吧
// 该函数带返回值，返回值为当前Tag
func (cur *GoTag) FuncR(fn func(*GoTag) *GoTag) *GoTag {
	return fn(cur)
}

// 在当前Tag下新建一个Tag，返回值为新建出来的Tag
func (cur *GoTag) Tag(name string, propertys ...P) *GoTag {
	newtag := NewTag()
	newtag.name = name
	if propertys != nil {
		newtag.property = make(P)
		for _, property := range propertys {
			for k, v := range property {
				newtag.property[k] = v
			}
		}
	}
	cur.Append(newtag)
	return newtag
}

// 设置当前Tag为self-Closing，返回当前Tag
func (cur *GoTag) SelfClosing() *GoTag {
	cur.selfclosing = true
	return cur
}

// 在当前Tag下新建一个Text，返回值为新建出来的Gotext（其实无意义）
func (cur *GoTag) Text(txt string) *GoText {
	newtext := Text(txt)
	cur.Append(newtext)
	return newtext
}
