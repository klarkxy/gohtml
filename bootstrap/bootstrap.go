/*
 * @Author       : Klarkxy
 * @Date         : 2020-02-18 17:48:45
 * @LastEditors  : Klarkxy
 * @LastEditTime : 2020-02-18 18:24:10
 * @FilePath     : \gohtml\bootstrap\bootstrap.go
 */
package bootstrap

import (
	"github.com/klarkxy/gohtml"
)

type GoBootstrap struct {
	gohtml.GoHtml
	bootstrapCSS *gohtml.GoTag // <link href="bootstrap.css" rel="stylesheet">
	jqueryJS     *gohtml.GoTag // <script src="jquery.js"></script>
	popperJS     *gohtml.GoTag // <script src="bootstrap.js"></script>
	bootstrapJS  *gohtml.GoTag // <script src="bootstrap.js"></script>
}

// 新建一个Bootstrap框架， 默认为Bootstrap4
func Bootstrap() *GoBootstrap {
	bootstrap := GoBootstrap{}
	bootstrap.Init()
	bootstrap.Html().Lang("zh-CN")
	// Meta部分
	bootstrap.Meta().Charset("utf-8")
	bootstrap.Meta().Http_equiv("X-UA-Compatible").Content("IE=edge")
	bootstrap.Meta().Name("viewport").Content("width=device-width, initial-scale=1")
	// Css引入
	bootstrap.SetBootstrapCSS("https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/css/bootstrap.min.css")
	// Js引入
	bootstrap.SetJqueryJS("https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js")
	bootstrap.SetPopperJS("https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js")
	bootstrap.SetBootstrapJS("https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/js/bootstrap.min.js")

	return &bootstrap
}

// 新建一个Bootstrap3框架
func Bootstrap3() *GoBootstrap {
	bootstrap := Bootstrap()
	// Css引入
	bootstrap.SetBootstrapCSS("https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css")
	// Js引入
	bootstrap.SetBootstrapJS("https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js")
	return bootstrap
}

// 初始化
func (cur *GoBootstrap) Init() *GoBootstrap {
	cur.GoHtml.Init()
	cur.bootstrapCSS = cur.Link()
	cur.jqueryJS = cur.Script()
	cur.popperJS = cur.Script()
	cur.bootstrapJS = cur.Script()
	return cur
}

// 设置bootstrap.css的URL
func (cur *GoBootstrap) SetBootstrapCSS(url string) *GoBootstrap {
	cur.bootstrapCSS.Href(url).Rel("stylesheet")
	return cur
}

// 设置jquery.js的URL
func (cur *GoBootstrap) SetJqueryJS(url string) *GoBootstrap {
	cur.jqueryJS.Src(url)
	return cur
}

// 设置bootstrap.js的URL
func (cur *GoBootstrap) SetBootstrapJS(url string) *GoBootstrap {
	cur.bootstrapJS.Src(url)
	return cur
}

// 设置bootstrap.js的URL
func (cur *GoBootstrap) SetPopperJS(url string) *GoBootstrap {
	cur.popperJS.Src(url)
	return cur
}
