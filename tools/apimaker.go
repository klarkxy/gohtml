/*
 * @Author       : Klarkxy
 * @Date         : 2020-02-16 14:04:28
 * @LastEditors  : Klarkxy
 * @LastEditTime : 2020-02-16 14:04:28
 * @FilePath     : \gohtml\tools\apimaker.go
 */
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

func foreach(srcFile, tsdFile, temp string) error {
	f, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	fmt.Println("Open", srcFile, "Success.")
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			return err
		}
		line = strings.TrimSpace(line)
		if line != "" {
			err = writeto(tsdFile, line, temp)
		}

		if err != nil {
			return err
		}
	}
	return nil
}
func writeto(filename, tagname, temp string) error {
	fileObj, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
		return err
	}
	content := maketag(tagname, temp)
	if _, err := io.WriteString(fileObj, content); err != nil {
		fmt.Println("Error appending to the file with os.OpenFile and io.WriteString.\n", content)
		return err
	}
	return nil
}
func newwrite(filename string) error {
	os.Remove(filename)
	fileObj, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
		return err
	}
	content := "package gohtml\n"
	if _, err := io.WriteString(fileObj, content); err != nil {
		fmt.Println("Error appending to the file with os.OpenFile and io.WriteString.\n", content)
		return err
	}
	return nil
}

type tagx struct {
	Old  string
	Low  string
	High string
}

func reply(temp string) string {
	return strings.ReplaceAll(temp, "-", "_")
}
func maketag(tagname, temp string) string {
	tag := tagx{
		Old:  tagname,
		Low:  strings.ToLower(reply(tagname)),
		High: Capitalize(strings.ToLower(reply(tagname))),
	}
	tmpl, err := template.New("tagx").Parse(temp)
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	err = tmpl.Execute(buffer, tag)
	if err != nil {
		panic(err)
	}
	//fmt.Println(buffer.String(), "\n")
	return buffer.String()
}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
func detal(name string, temp string) {

	inFile := name + ".txt"
	outFile := "../" + name + ".go"
	newwrite(outFile)
	err := foreach(inFile, outFile, temp)
	if err != nil {
		fmt.Println("Err:", err)
	}
}

func main() {
	detal("api", api)
	detal("self", self)
	detal("attr", attr)
	detal("singleattr", singleattr)
}

var api = `
//  <{{.Old}}></{{.Old}}>
func {{.High}}() *GoTag {
	return Tag("{{.Low}}")
}
func (cur *GoTag) {{.High}}() *GoTag {
	return cur.Tag("{{.Low}}")
}
`
var self = `
//  <{{.Old}}>
func {{.High}}() *GoTag {
	return Tag("{{.Low}}").Self()
}
func (cur *GoTag) {{.High}}() *GoTag {
	return cur.Tag("{{.Low}}").Self()
}
`
var attr = `
//  <?? {{.Old}}="{val}">
func (cur *GoTag) {{.High}}(val string) *GoTag {
	return cur.Attr("{{.Low}}", val)
}
`
var singleattr = `
//  <?? {{.Old}}>
func (cur *GoTag) {{.High}}() *GoTag {
	return cur.SingleAttr("{{.Low}}")
}
`
