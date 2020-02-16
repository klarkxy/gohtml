# GoHtml
Version : 1.0
使用go动态生成html。

`go get gitee.com/Klarkxy/gohtml`

[toc]


-----------------
## 功能
* **GoTag**
  GoTag代表一个Html的Tag。使用String()接口来产生形如`<{name} {attr}>{tags}</{name}>`的html代码。
  GoTag中可以嵌入多个GoTag，具体看下面例子。
* **GoHtml**
  @TODO


## GoTag部分
使用`gohtml.Tag(name)`来创建一个GoTag。
GoTag的主要接口：
* **String**()
  将GoTag转化为html输出。
  **Return: string**
* **Append**(tags ...fmt.Stringer)
  往当前GoTag中增加若干条{tag}，可以添加任意定义了`String()`接口的类型，在输出时将会调用此接口将其转化为字符串
  **Return**: *GoTag  **当前**GoTag
* **Tag**(name string)
  往当前GoTag中新增一条GoTag。
  **Return**: *GoTag  **新增**GoTag
* **Text**(str string)
  往当前GoTag中新增一段string。
  **Return**: *GoTag  **当前**GoTag
* **Attr**(name, val string)
  往当前GoTag中新增一条属性`name="val"`。
  **Return**: *GoTag  **当前**GoTag
* **Self**()
  将当前GoTag设置为SelfClosing-Tag，只包含{name}和{attr}t
  将会产生形如`<{name} {attr}>`的html代码
  注意，对设置了Self()的GoTag，所有Append()、Tag()、Text()等增加的{tags}都会失效。
  **Return**: *GoTag  **当前**GoTag

### 第一个例子
```go
htm1 := gohtml.Tag("p").Text("Hello World!")
fmt.Println(htm.String())
```
输出
```html
<p >Hello World! </p>
```
<p >Hello World! </p>


### 设置属性
#### 设置一个属性
```go
htm := gohtml.Tag("a").Attr("href", "https://www.zhihu.com/").Text("打开知乎")
fmt.Println(htm.String())
```
输出
```html
<a href="https://www.zhihu.com/" >打开知乎 </a>
```
<a href="https://www.zhihu.com/" >打开知乎 </a>


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
<button type="button" href="https://www.zhihu.com/" >打开知乎 </button>
<button type="button" href="https://www.zhihu.com/" >打开知乎 </button>
```
<button type="button" href="https://www.zhihu.com/" >打开知乎 </button>


### SelfClosing Tag
html的SelfClosing-Tag
```go
htm1 := gohtml.Tag("meta").Attr("http-equiv", "X-UA-Compatible").Attr("content", "IE=edge").Self()
fmt.Println(htm1.String())
```
输出
```html
<meta http-equiv="X-UA-Compatible" content="IE=edge" >
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
<table ><tr ><th >1 </th> <th >2 </th> </tr> <tr ><th >3 </th> <th >4 </th> </tr> </table>
```
<table ><tr ><th >1 </th> <th >2 </th> </tr> <tr ><th >3 </th> <th >4 </th> </tr> </table>

### 更为优美的书写方式
目前已经支持大部分html-tag，可以直接进行调用。
同时支持`gohtml.Xxx()`或者`GoTag.Xxx()`的方式。效果同`gohtml.Tag("xxx")`或者`GoTag.Tag("xxx")`
同样，也支持了`GoTag.Yyy(val string)`的方式，直接设置属性，效果同`GoTag.Attr("yyy", val)`
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
<form><h2>Please Sign In</h2><input type="username" placeholder="Username"><input type="password" placeholder="Password"><div><label><input type="checkbox" value="Remember me">Remember me.</label></div><button type="submit">Sign in</button></form>
```
<form><h2>Please Sign In</h2><input type="username" placeholder="Username"><input type="password" placeholder="Password"><div><label><input type="checkbox" value="Remember me">Remember me.</label></div><button type="submit">Sign in</button></form>


-----------------

## 联系我
* QQ: 278370456
* Mail: 278370456@qq.com
如果您在使用时遇见了问题，或者有意见、建议，请务必联系我，谢谢！

以上