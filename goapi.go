package gohtml

// 此文件设置 快捷语义

// 设置属性 name="value"，返回值为当前GoTag
// 一次设置一组，如有重复则覆盖
func (cur *GoTag) P(name, value string) *GoTag {
	cur.property[name] = value
	return cur
}

// 设置属性 name="name"，返回值为当前GoTag
// 一次设置一组，如有重复则覆盖
func (cur *GoTag) PS(name string) *GoTag {
	cur.property[name] = name
	return cur
}

// if表达式，根据条件判断设置属性
// IfP(flag, A, B) ==> if (flag) A="B" else nil
// IfP(flag, A, B, C, D) ==> if (flag) A="B" else C="D"
func (cur *GoTag) IfP(flag bool, properties ...string) *GoTag {
	if properties != nil {
		if len(properties) >= 4 {
			if !flag {
				cur.P(properties[2], properties[3])
			}
		}
		if len(properties) >= 2 {
			if flag {
				cur.P(properties[0], properties[1])
			}
		}
	}
	return cur
}

// 往cur中插入一个Tag，返回值为插入的这个Tag
// 可以直接一次性导入多个属性，按顺序分为一对一对
func (cur *GoTag) T(name string, properties ...string) *GoTag {
	ths := cur.Tag(name)
	if properties != nil {
		for i := 0; i < len(properties); i += 2 {
			if i+1 >= len(properties) {
				// 参数不成对，直接跳过最后的单个
				break
			}
			ths.P(properties[i], properties[i+1])
		}
	}
	return ths
}

// 往cur中插入一个Self-Closing Tag，返回值为插入的这个Tag
// 可以直接一次性导入多个属性，按顺序分为一对一对
func (cur *GoTag) S(name string, properties ...string) *GoTag {
	ths := cur.Tag(name)
	if properties != nil {
		for i := 0; i < len(properties); i += 2 {
			if i+1 >= len(properties) {
				// 参数不成对，直接跳过最后的单个
				break
			}
			ths.P(properties[i], properties[i+1])
		}
	}
	return ths.SelfClosing()
}
