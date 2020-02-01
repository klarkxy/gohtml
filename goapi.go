package gohtml

// 此文件设置 快捷语义

// 设置属性，返回值为当前GoTag
// 一次设置一组，如有重复则覆盖
func (cur *GoTag) P(name, value string) *GoTag {
	cur.property[name] = value
	return cur
}

// 往cur中插入一个Tag，返回值为插入的这个Tag
func (cur *GoTag) T(name string) *GoTag {
	return cur.Tag(name)
}

// 往cur中插入一个Self-Closing Tag，返回值为插入的这个Tag
func (cur *GoTag) S(name string) *GoTag {
	return cur.Tag(name).SelfClosing()
}

// 返回上一层
func (cur *GoTag) M() *GoTag {
	return cur.Master()
}
