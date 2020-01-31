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
其中`gohtml.P`==`map[string]string`，可以设置多个

##### 嵌套