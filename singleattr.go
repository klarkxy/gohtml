package gohtml

//  <?? required>
func (cur *GoTag) Required() *GoTag {
	return cur.SingleAttr("required")
}

//  <?? autofocus>
func (cur *GoTag) Autofocus() *GoTag {
	return cur.SingleAttr("autofocus")
}
