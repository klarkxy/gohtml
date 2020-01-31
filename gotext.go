package gohtml

type GoText struct {
	txt string
}

func (cur *GoText) String() string {
	return cur.txt
}
