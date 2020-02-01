# gohtml
一个使用go动态生成html的包，通常用于后端渲染html。

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
<table border="1" ><tr ><td >1 </td> <td >2 </td> </tr> <tr ><td >3 </td> <td >4 </td> </tr> </table>
```
<table border="1" ><tr ><td >1 </td> <td >2 </td> </tr> <tr ><td >3 </td> <td >4 </td> </tr> </table>
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
<form ><input type="text" placeholder="用户名" > <input type="password" placeholder="密码" > <button type="submit" >登录 </button> </form>
```
<form ><input type="text" placeholder="用户名" > <input type="password" placeholder="密码" > <button type="submit" >登录 </button> </form>

### Html模板
使用`tag.Html(name, filepath)`来导入一个指定位置的Html
如果要使用模板，则使用`tag.Html(name, filepath).Excute(data)`来导入模板。模板使用方式详见`html/template`
```go
htm5 := gohtml.NewTag()
PWD, _ := os.Getwd()
htm5.Html("index.html", PWD+"/index.html").Execute("Hello World!")
fmt.Println(htm5.String())
```
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8"/>
    <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
    <meta name="viewport" content="width=device-width, initial-scale=1"/>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css" rel="stylesheet"/>
  </head>

  <body>
    <div class="container">
      <form class="form-signin" method="POST" action="/">
        <h2 class="form-signin-heading">Hello World!</h2>
        <label for="inputUsername" class="sr-only">用户名</label>
        <input type="username" name="username" id="inputUsername" class="form-control" placeholder="用户名" required autofocus>       
        <label for="inputPassword" class="sr-only">密码</label>
        <input type="password" name="password" id="inputPassword" class="form-control" placeholder="密码" required>
        <button class="btn btn-lg btn-primary btn-block" type="submit">登录</button>
      </form>
    </div>
    <script src="https://cdn.jsdelivr.net/npm/jquery@1.12.4/dist/jquery.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js"></script>
  </body>
</html>



-----------
## 联系我
* QQ: 278370456
* Mail: 278370456@qq.com

如果您在使用时遇见了问题，或者有什么意见或者建议，请务必联系我，谢谢。