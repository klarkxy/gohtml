package gohtml

// 作为参数使用，能少写很多字符
type P map[string]string

// 包含两种类型：GoText/GoTag
type gotext interface {
	text() string
}

// 创建一个GoTag对象
func NewTag() *GoTag {
	return &GoTag{}
}

// 创建一个GoText对象
func Text(txt string) *GoText {
	return &GoText{txt: txt}
}

func Tag(name string, propertys ...P) *GoTag {

	cur := NewTag()
	cur.name = name
	for _, property := range propertys {
		for k, v := range property {
			cur.property[k] = v
		}
	}
	return cur
}
