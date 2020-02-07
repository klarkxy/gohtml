# gohtml
一个使用go动态生成html的包，通常用于后端渲染html。

另一个README：
[gohtml：使用go进行html开发](https://zhuanlan.zhihu.com/p/104894925)


[toc]

## 安装教程

`go get gitee.com/Klarkxy/gohtml`

## 使用说明

### 第一个例子
```go
htm1 := gohtml.NewTag()
htm1.Tag("p").Text("Hello World!")
fmt.Println(htm1.String())
```
输出
```html
<p >Hello World! </p>
```
<p >Hello World! </p>


### 加入一些属性
```go
htm2 := gohtml.NewTag()
htm2.Tag("a", gohtml.P{"href": "www.baidu.com"}).Text("前往百度")
fmt.Println(htm2.String())
```
输出
```html
<a href="www.baidu.com" >前往百度 </a>
```
<a href="www.baidu.com" >前往百度 </a>


其中`gohtml.P`=`map[string]string`，可以通过导入map来导入多个属性

### 嵌套
```go
htm3 := gohtml.NewTag()
htm3.Tag("table", gohtml.P{"border": "1"}).Func(func(Htm *gohtml.GoTag) {
	Htm.Tag("tr").Tag("td").Text("1").
		Master().Tag("td").Text("2")
	Htm.Tag("tr").Tag("td").Text("3").
		Master().Tag("td").Text("4")
	})
```
输出
```html
<table border="1" >
<tr >
<td >1</td>
<td >2</td>
</tr>
<tr >
<td >3</td>
<td >4</td>
</tr>
</table>
```
<table border="1" >
<tr >
<td >1</td>
<td >2</td>
</tr>
<tr >
<td >3</td>
<td >4</td>
</tr>
</table>


该例中使用了两种嵌套方式，一种是使用`Func(func(*gohtml.GoTag)`函数进行嵌套，另一种是使用`Master()`函数返回上一层

### 快捷语义

* `T(name string)`
    同Tag(name)，插入一个Tag，返回这个Tag
* `S(name)`
    插入一个Self-Closing Tag
* `P(param, value string)`
    往当前Tag中加入一条属性`param="value"`

一般通常情况下建议使用这种模式
```go
htm4 := gohtml.NewTag()
form := htm4.T("form")
form.S("input").P("type", "text").P("placeholder", "用户名")
form.S("input").P("type", "password").P("placeholder", "密码")
form.T("button").P("type", "submit").Text("登录")
fmt.Println(htm4.String())
```
输出
```html
<form >
<input placeholder="用户名" type="text" >
<input type="password" placeholder="密码" >
<button type="submit" >登录</button>
</form>
```


同时，`S()`和`T()`中也可以直接成对插入属性
```go
	htm6 := gohtml.NewTag()
	htm6.S("input", "type", "text", "placeholder", "Hello World!")
	htm6.T("button", "type", "submit").Text("登录")
	fmt.Println(htm6.String())
```
输出
```html
<input type="text" placeholder="Hello World!" >
<button type="submit" >登录</button>
```

### IfP语义
根据条件设置属性
* IfP(flag, A, B) ==> if (flag) A="B" else nil
* IfP(flag, A, B, C, D) ==> if (flag) A="B" else C="D"
```go
	htm7 := gohtml.NewTag()
	canedit := false
	htm7.S("input", "type", "text").IfP(canedit, "readonly", "readonly")
	canedit = true
	htm7.S("input", "type", "text").IfP(canedit, "readonly", "readonly")
	ispassword := false
	htm7.S("input").IfP(ispassword, "type", "password", "type", "text")
	ispassword = true
	htm7.S("input").IfP(ispassword, "type", "password", "type", "text")
	fmt.Println(htm7.String())
```
输出
```html
<input type="text" >
<input type="text" readonly="readonly" >
<input type="text" >
<input type="password" >
```


### Html模板
使用`tag.Html(name, filepath)`来导入一个指定位置的Html
如果要使用模板，则使用`tag.Html(name, filepath).Excute(data)`来导入模板。模板使用方式详见`html/template`
```go
htm5 := gohtml.NewTag()
PWD, _ := os.Getwd()
htm5.Html("index.html", PWD+"/index.html").Execute("Hello World!")
fmt.Println(htm5.String())
```

-----------
## 联系我
* QQ: 278370456
* Mail: 278370456@qq.com

如果您在使用时遇见了问题，或者有什么意见或者建议，请务必联系我，谢谢。~~~