package gohtml

// 作为参数使用，能少写很多字符
type P map[string]string

// 创建一个GoTag对象
func NewTag() *GoTag {
	return &GoTag{}
}

func Tag(name string, propertys ...P) *GoTag {
	cur := NewTag()
	cur.name = name
	if propertys != nil {
		cur.property = make(P)
		for _, property := range propertys {
			for k, v := range property {
				cur.property[k] = v
			}
		}
	}
	return cur
}
