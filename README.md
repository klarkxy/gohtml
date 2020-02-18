# GoHtml
Version : 1.0
使用go动态生成html。

`go get github.com/klarkxy/gohtml`


-----------------
## 功能
* **GoTag**
  GoTag代表一个Html的Tag。使用`String()`接口来产生形如`<{name} {attr}>{tags}</{name}>`的html代码。
* **GoHtml**
  GoHtml代表一个Html的页面。使用`String()`接口来生成Html字符串，或者使用`MinString()`接口来生成紧凑的Html字符串


## GoTag部分
使用`gohtml.Tag(name)`来创建一个GoTag。
### 主要接口：
```go
// 将GoTag转化为html字符串输出
func (cur *GoTag) String() string;
// 在当前Tag中插入若干个Tag/Text，返回当前Tag
func (cur *GoTag) Append(tags ...fmt.Stringer) *GoTag;
// 在当前Tag下新建一个Tag，返回新建出来的Tag
func (cur *GoTag) Tag(name string) *GoTag;
// 设置当前Tag为self-Closing，返回当前Tag
func (cur *GoTag) Self() *GoTag;
// 在当前Tag中插入一条string，返回当前Tag
func (cur *GoTag) Text(str string) *GoTag;
// 设置属性，返回当前Tag
func (cur *GoTag) Attr(attrs ...string) *GoTag;
```
### 第一个例子
```go
htm1 := gohtml.Tag("p").Text("Hello World!")
fmt.Println(htm.String())
```
输出
```html
<p>
  Hello World!
</p>
```


### 设置属性
#### 设置一个属性
```go
htm := gohtml.Tag("a").Attr("href", "https://www.zhihu.com/").Text("打开知乎")
fmt.Println(htm.String())
```
输出
```html
<a href="https://www.zhihu.com/">
  打开知乎
</a>
```


#### 设置多个属性
可以多次调用Attr接口，也可以一次性调用中传入多个参数。
```go
htm1 := gohtml.Tag("button").Attr("type", "button").Attr("href", "https://www.zhihu.com/").Text("打开知乎")
htm2 := gohtml.Tag("button").Attr("type", "button", "href", "https://www.zhihu.com/").Text("打开知乎")
fmt.Println(htm1.String())
fmt.Println(htm2.String())
```
输出
```html
<button type="button" href="https://www.zhihu.com/">
  打开知乎
</button>
<button type="button" href="https://www.zhihu.com/">
  打开知乎
</button>
```


### SelfClosing Tag
html的SelfClosing-Tag
```go
htm1 := gohtml.Tag("meta").Attr("http-equiv", "X-UA-Compatible").Attr("content", "IE=edge").Self()
fmt.Println(htm1.String())
```
输出
```html
<meta http-equiv="X-UA-Compatible" content="IE=edge">
```


### GoTag的嵌套
```go
table := gohtml.Tag("table")
tr1 := table.Tag("tr")
tr1.Tag("th").Text("1")
tr1.Tag("th").Text("2")
tr2 := table.Tag("tr")
tr2.Tag("th").Text("3")
tr2.Tag("th").Text("4")
fmt.Println(table.String())
```
输出
```html
<table>
  <tr>
    <th>
      1
    </th>
    <th>
      2
    </th>
  </tr>
  <tr>
    <th>
      3
    </th>
    <th>
      4
    </th>
  </tr>
</table>
```

### 更为优美的书写方式
目前已经支持大部分html-tag，可以直接进行调用。  

同时支持`gohtml.Xxx()`或者`GoTag.Xxx()`的方式。效果同`gohtml.Tag("xxx")`或者`GoTag.Tag("xxx")`  

同样，也支持了`GoTag.Yyy(val string)`的方式，直接设置属性，效果同`GoTag.Attr("yyy", val)`  

详见[api.go](./api.go)、[attr.go](./attr.go)、[self.go](./self.go)、[singleattr.go](./singleattr.go)

由于html的属性种类繁多，可能会有缺漏，请务必提醒我。
```go
form := gohtml.Form()
form.H2().Text("Please Sign In")
form.Input().Type("username").Placeholder("Username").Required().Autofocus()
form.Input().Type("password").Placeholder("Password").Required()
divlabel := form.Div().Label()
divlabel.Input().Type("checkbox").Value("Remember me")
divlabel.Text("Remember me.")
form.Button().Type("submit").Text("Sign in")
fmt.Println(form.String())
```
输出
```html
<form>
  <h2>
    Please Sign In
  </h2>
  <input type="username" placeholder="Username">
  <input type="password" placeholder="Password">
  <div>
    <label>
      <input type="checkbox" value="Remember me">
      Remember me.
    </label>
  </div>
  <button type="submit">
    Sign in
  </button>
</form>
```

## GoHtml部分
通常意义上，我们可以把一个html文件分为以下6个部分。
```html
<!DOCTYPE html>
<html>
  <head>
      {{.Meta}}
      {{.Link}}
      {{.HtmlHead}}
  </head>
  <body>
      {{.Body}}
      {{.Script}}
  </body>
  {{.Html}}
</html>
```
我们调用`gohtml.NewHtml()`来生成一个GoHtml，调用其String()接口来将之转换成html字符串。  

### 常用接口
```go
// 生成一个GoHtml
func NewHtml() *GoHtml;
// 转化为html字符串输出。如果不希望有空格和换行，使用MinString()
func (cur *GoHtml) String() string;
func (cur *GoHtml) MinString() string;
// 获取对应Tag
func (cur *GoHtml) Html() *GoTag;
func (cur *GoHtml) Head() *GoTag;
func (cur *GoHtml) Body() *GoTag;
// 增加一条Tag
func (cur *GoHtml) Meta() *GoTag;
func (cur *GoHtml) Link() *GoTag;
func (cur *GoHtml) Script() *GoTag;
```
* 对于`Html()`接口，调用`Tag()`接口增加的Tag将会增加在`{{.Html}}`处
* 对于`Head()`接口和`Body()`接口，任何增加属性的行为将会失效
* 对于`Meta()`、`Link()`、`Script()`接口，调用接口会增加一条对应的Tag

### 一个空的GoHtml
```go
htm := gohtml.NewHtml()
fmt.Println("String:")
fmt.Println(htm.String())
fmt.Println("MinString:")
fmt.Println(htm.MinString())
```
输出
```html
String:
<!DOCTYPE html>
<html>
  <head>
  </head>
  <body>
  </body>
</html>
MinString:
<!DOCTYPE html><html><head></head><body></body></html>
```
### 一个基本的Bootstrap3框架
```go
bootstrap := gohtml.NewHtml()
bootstrap.Html().Lang("zh-CN")
// Meta部分
bootstrap.Meta().Charset("utf-8")
bootstrap.Meta().Http_equiv("X-UA-Compatible").Content("IE=edge")
bootstrap.Meta().Name("viewport").Content("width=device-width, initial-scale=1")
// Css引入
bootstrap.Link().Href("https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css").Rel("stylesheet")
// Js引入
bootstrap.Script().Src("https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js")
bootstrap.Script().Src("https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js")
// Head
bootstrap.Head().Title().Text("Bootstrap 101 Template")
// Body
bootstrap.Body().H1().Text("Hello World!")

fmt.Println(bootstrap.String())
```
输出
```html
<!DOCTYPE html>
<html lang="zh-CN">
  <head>

    <meta charset="utf-8">
    <meta http_equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css">

    <title>
      Bootstrap 101 Template
    </title>
  </head>
  <body>

    <h1>
      Hello World!
    </h1>

    <script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js">    </script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js">    </script>
  </body>
</html>
```
现在已经加入了Bootstrap模块，使用`import "github.com/klarkxy/gohtml/bootstrap`来直接使用吧！
[github.com/klarkxy/gohtml/bootstrap](./bootstrap/README.md)


-----------------

## 联系我
* QQ: 278370456
* Mail: 278370456@qq.com
如果您在使用时遇见了问题，或者有意见、建议，请务必联系我，谢谢！  


以上