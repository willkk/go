# Go实战

## 1. go 命令(build, install, test，快速略过)

## 2. cgo(调用C代码)  
import "C"之前存放被注释的C代码。详细参考https://golang.org/cmd/cgo/。使用cgo命令添加C编译选项，如// #cgo LDFLAGS: -L/go/src/foo/libs -lfoo。  
## 3. vendor(vendor目录)  
将项目使用的第三方库放到vendor目录，防止项目间不同版本间冲突。解决第三方库多项目多版本冲突问题。并不解决第三方库版本管理和自动更新问题，使用govendor库。  
## 4. 配置文件(json config)  
配置文件struct定义：  
![](https://github.com/willkk/go/blob/master/images/goskill_jsonconf.png)  
**说明：可选字段可以声明成指针类型，json在解析时，如果没有设置该字段，指针为nil。同样可以用在json消息可选字段的场景。**  
配置文件，如server.conf：  
![](https://github.com/willkk/go/blob/master/images/goskill_jsonfile.png)  
使用时：  
![](https://github.com/willkk/go/blob/master/images/goskill_jsonparse.png)

## 5. 函数：返回命名参数  
对返回参数命名和传入参数命名是一样的，都会生成对应类型的空对象，比如int为0，string为""，map，chan和pointer为nil。恰当使用可以节省大量冗余代码。  

普通代码：  
![](https://github.com/willkk/go/blob/master/images/goskill_func.png)
![](https://github.com/willkk/go/blob/master/images/goskill_func2.png) 

使用命名返回参数：  
![](https://github.com/willkk/go/blob/master/images/goskill_func_new.png)
![](https://github.com/willkk/go/blob/master/images/goskill_func_new2.png)

## 6. defer函数  
恰当使用defer函数，减少冗余代码。

普通代码：  
![](https://github.com/willkk/go/blob/master/images/goskill_nodefer.png)  
defer:  
![](https://github.com/willkk/go/blob/master/images/goskill_defer.png)

## 7. log库(seelog)  
#### 第三方库seelog:   
支持同步/异步（获取频率：循环，定时，自适应）；支持丰富的输出级别，格式，终端和颜色；支持不同级别log的文件过滤；支持log文件回滚；支持动态更新logger设置。  
#### 实现自定义格式log：  
runtime.Caller(1)返回调用函数所在的全文件路径名和行号；添加前缀log.SetPrefix(); os.Getpid()获取PID等。 

好的log格式举例（SRS）：  
![](https://github.com/willkk/go/blob/master/images/goskill_srslog.png)

**说明：基本信息包括标识pid和tid，错误级别，错误号等。需要唯一标识每一个request，比如请求的userid，协程的id，线程id等。快速查看单个请求的全流程结果，方便定位功能级和业务级错误。错误号定义：分类定义，如http 1xx-5xx。**

## 8. web framework(beego, iris，martini等)  
beego:使用MVC+ORM+RESTful，组件齐全，使用方便，适合简单的中小型web server，框架比较臃肿，实现性能也不并高。  
Iris: 速度最快的web framework，简单。 

MVC(应用分层)：  
M：Model，数据模型，一般映射数据库表，作为数据层；V：View，视图，一般通过template生成，页面上有占位符，后端代码只需要填充数据到占位符即可，作为表示层；C：Controller，控制器，连接页面View和后端数据库Model，实现控制逻辑，作为逻辑层。  

|   View    |
|-----------|
| Controller|
|   Model   |
| DataBase  |

