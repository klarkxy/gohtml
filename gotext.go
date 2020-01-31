package gohtml

type GoText struct {
	txt string
}

func (cur *GoText) text() string {
	return cur.txt
}
