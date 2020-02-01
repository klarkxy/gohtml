# gohtml
[tag]
#### 介绍
一个使用go动态生成html的包，通常用于后端渲染html。


#### 安装教程

`go get gitee.com/Klarkxy/gohtml`

#### 使用说明

##### 第一个例子
```
htm1 := gohtml.NewTag()
htm1.Tag("p").Text("Hello World!")
fmt.Println(htm1.String())
```
输出
```
<p >Hello World! </p>
```
<p >Hello World! </p>


##### 加入一些属性
```
htm2 := gohtml.NewTag()
htm2.Tag("a", gohtml.P{"href": "www.baidu.com"}).Text("前往百度")
fmt.Println(htm2.String())
```
输出
```
<a href="www.baidu.com" >前往百度 </a>
```
<a href="www.baidu.com" >前往百度 </a>


其中`gohtml.P`==`map[string]string`，可以设置多个

##### 嵌套
```
htm3 := gohtml.NewTag()
htm3.Tag("table", gohtml.P{"border": "1"}).Func(func(Htm *gohtml.GoTag) {
	Htm.Tag("tr").Tag("td").Text("1").
		Master().Tag("td").Text("2")
	Htm.Tag("tr").Tag("td").Text("3").
		Master().Tag("td").Text("4")
	})
```
输出
```
<table border="1" ><tr ><td >1 </td> <td >2 </td> </tr> <tr ><td >3 </td> <td >4 </td> </tr> </table>
```
<table border="1" ><tr ><td >1 </td> <td >2 </td> </tr> <tr ><td >3 </td> <td >4 </td> </tr> </table>
该例中使用了两种嵌套方式，一种是使用`Func(func(*gohtml.GoTag)`函数进行嵌套，另一种是使用`Master()`函数返回上一层

#### 快捷语义

* `T(name string)`
    同Tag(name)，插入一个Tag，返回这个Tag
* `S(name)`
    插入一个Self-Closing Tag
* `P(param, value string)`
    往当前Tag中加入一条属性`param="value"`
* `M()`
    获取上一层Tag

### 联系我
* QQ: 278370456
* Mail: 278370456@qq.com

如果您在使用时遇见了问题，或者有什么意见或者建议，请务必联系我，谢谢。