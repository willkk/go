# Go实战

## 1. go 命令(build, install, test，快速略过)

## 2. cgo(调用C代码)

import "C"之前存放被注释的C代码。详细参考https://golang.org/cmd/cgo/。使用cgo命令添加C编译选项，如// #cgo LDFLAGS: -L/go/src/foo/libs -lfoo。

## 3. vendor(vendor目录)

将项目使用的第三方库放到vendor目录，防止项目间不同版本间冲突。解决第三方库多项目多版本冲突问题。并不解决第三方库版本管理和自动更新问题，使用govendor库。

## 4. 配置文件(json config)

配置文件struct定义：

![](https://github.com/willkk/go/blob/master/images/goskill_jsonconf.png)

#### 说明：可选字段可以声明成指针类型，json在解析时，如果没有设置该字段，指针为nil。同样可以用在json消息可选字段的场景。

配置文件，如server.conf：

![](https://github.com/willkk/go/blob/master/images/goskill_jsonfile.png)

使用时：

![](https://github.com/willkk/go/blob/master/images/goskill_jsonparse.png)

## 5. log库(seelog)

第三方库seelog: 支持同步/异步（获取频率：循环，定时，自适应）；支持丰富的输出级别，格式，终端和颜色；支持不同级别log的文件过滤；支持log文件回滚；支持动态更新logger设置。

实现自定义格式log：runtime.Caller(1)返回调用函数所在的全文件路径名和行号；添加前缀log.SetPrefix(); os.Getpid()获取PID等。

## 6. web framework(beego, iris，martini等)

beego使用MVC+ORM+RESTful，组件齐全，使用方便，适合简单的中小型web server，框架比较臃肿，实现性能也不并高。

MVC：主要做web界面，不细讲。M：Model，数据模型，一般映射数据库表；V：View，视图，一般通过template生成，页面上有占位符，后端代码只需要填充数据到占位符即可；C：Controller，控制器，连接页面View和后端数据库Model，实现控制逻辑。
