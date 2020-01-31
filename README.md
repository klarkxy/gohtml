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