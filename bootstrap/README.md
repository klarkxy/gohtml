# Bootstrap 模块
安装：
`go get github.com/klarkxy/gohtml/bootstrap`


### 快速开始
```go
bootstrap := bootstrap.Bootstrap()
fmt.Println(bootstrap.String())
```
输出
```html
<!DOCTYPE html>
<html lang="zh-CN">
  <head>
    <meta charset="utf-8">
    <meta content="IE=edge" http_equiv="X-UA-Compatible">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/css/bootstrap.min.css" rel="stylesheet">
  </head>
  <body>
    <script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js">    </script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js">    </script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/js/bootstrap.min.js">    </script>
  </body>
</html>
```
这是一个仅包含Bootstrap框架的GoHtml对象。

当前默认Bootstrap版本为4.4.1，如果要创建一个Bootstrap3的框架，请使用`bootstrap := bootstrap.Bootstrap3()`，默认版本为3.3.7

### 主要接口

```go
// 生成一个包含Bootstrap框架的GoHtml对象
func Bootstrap() *GoBootstrap;
func Bootstrap3() *GoBootstrap;
// 转化为html字符串输出。如果不希望有空格和换行，使用MinString()
func (cur *GoBootstrap) String() string;
func (cur *GoBootstrap) MinString() string;
// 获取对应Tag
func (cur *GoBootstrap) Html() *GoTag;
func (cur *GoBootstrap) Head() *GoTag;
func (cur *GoBootstrap) Body() *GoTag;
// 增加一条相应的Tag
func (cur *GoBootstrap) Meta() *GoTag;
func (cur *GoBootstrap) Link() *GoTag;
func (cur *GoBootstrap) Script() *GoTag;
// 更换默认CSS、JS的地址
func (cur *GoBootstrap) SetBootstrapCSS(url string) *GoBootstrap;
func (cur *GoBootstrap) SetJqueryJS(url string) *GoBootstrap;
func (cur *GoBootstrap) SetBootstrapJS(url string) *GoBootstrap;
func (cur *GoBootstrap) SetPopperJS(url string) *GoBootstrap;
```

### 其他组件
@TODO