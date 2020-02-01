package gohtml

import (
	"bytes"
	"fmt"
	"html/template"
)

type GoHtm struct {
	name     string
	filepath string
	data     interface{}
	master   *GoTag
}

func (cur *GoHtm) String() string {
	tmpl := template.Must(template.New(cur.name).ParseFiles(cur.filepath))
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, cur.data)
	if err != nil {
		fmt.Println(err)
	}

	return buf.String()
}

// 设置模板参数。 模板用法详见html/template
func (cur *GoHtm) Execute(data interface{}) *GoHtm {
	cur.data = data
	return cur
}

func (cur *GoHtm) Master() *GoTag {
	if cur.master == nil {
		return &GoTag{}
	} else {
		return cur.master
	}
}
