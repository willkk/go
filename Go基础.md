# Go语言基础

## 1.基础语法

#### a)Go变量、结构体和函数声明及定义


变量声明

![](https://github.com/willkk/go/blob/master/images/gobase_var.png)

结构体声明

![](https://github.com/willkk/go/blob/master/images/gobase_typedef.png)

函数声明

![](https://github.com/willkk/go/blob/master/images/gobase_func.png)

#### b)语句（for, go, defer, select, switch）

for语句

![](https://github.com/willkk/go/blob/master/images/gobase_for.png)

defer语句

![](https://github.com/willkk/go/blob/master/images/gobase_defer.png)

switch语句

![](https://github.com/willkk/go/blob/master/images/gobase_switch.png)

select语句

![](https://github.com/willkk/go/blob/master/images/gobase_select.png)

#### c)包管理

代码行宽：建议80-120个字符宽度（go语言是80字符宽度）。

![](https://github.com/willkk/go/blob/master/images/gobase_linewidth.png)

包内变量和函数：Go采用大小驼峰命名法来命名变量和函数，实现其显隐性(下划线适合过程函数，驼峰适合对象)；建议就近声明变量、常量和函数（局部化，模块化）。

![](https://github.com/willkk/go/blob/master/images/gobase_camel.png)

包初始化：在包内每个文件提供func init() {}来完成该包内全部的初始化工作。

![](https://github.com/willkk/go/blob/master/images/gobase_initfunc.png)

包文件及目录命名：目录及文件名称小写，尽量简短；必要时可用下划线。

![](https://github.com/willkk/go/blob/master/images/gobase_filedir.png)

包目录结构：各子模块如果可以一个go源文件完成就不再建立目录；否则建立独立子目录存放该模块的多个文件。包级公共模块放在包顶级目录。

![](https://github.com/willkk/go/blob/master/images/gobase_pkgsubdir.png)

## 2.数据类型

#### a) channel(内存队列, 读/写/读写)

FIFO同步内存队列，实现协程间数据通信。所以，内存效率高，多用，常用。另外，还可以用于goroutine的退出控制，close(chan)会立即让全部的管道接收操作返回。

#### b) map(hash, 无序)

实现结构是hash，成员无序。空间换时间，内存占用较多，仔细考虑。go里面对应的tree实现是treemap库。

#### c) struct(嵌入/提升/继承, 覆盖)

struct匿名字段也叫嵌入字段，类似于C++的继承，字段名称为该类型名称。go会把嵌入字段的成员提升到当前结构体作用域，效果等效于C++中公有继承时访问基类成员。基类成员类型可以是结构体类型或其指针类型。在声明该结构体并同时初始化时，需要嵌套构造，其他地方可以直接访问基类字段。结构体不能递归嵌套，不管直接或间接。外层字段会覆盖内层相同名称字段。

规则：1). 嵌入字段类型为结构体类型，子类类型可以访问基类或基类指针的方法集合；子类指针可以访问基类指针的方法集合。

2). 嵌入字段类型为结构体指针，子类类型及子类指针都可以访问基类及基类指针的方法集合。

#### d) interface和method(接口类型, 嵌套/继承)

接口类型包括一个方法集合，接口和实现类型之间不要求显式继承，只需要该类型或其指针实现该接口定义的方法集合即可。任意类型都实现空接口对象interface{}。一个类型可以同时实现多个接口。接口类型也可以像结构体一样实现嵌入/继承，递归除外。

method可以定义在任何类型上，除了指针和interface接口。比如，可以在一个函数上定义一个方法。参考http包HandlerFunc（类似只包含一个方法的匿名interface）。

#### e) func(过程函数, 成员函数, 参数名)

（参考函数定义部分）

## 3.并发模型

Do not communicate by sharing memory; instead, share memory by communicating.（不要通过共享来通信，而是通过通信来共享。）

开启指定数量的goroutine等待在每个channel上，并控制goroutine的退出。channel是协程安全的，同步的。

goroutine一般用途：

1) 每个请求或客户开启一个goroutine（比如net/http包server.go）。

2) 提前开启固定数量的goroutine，等待在几个channel上，循环处理请求。

channel一般用途：

1) channel作为数据队列，用于排队请求和返回消息。

2) channel作为信号管道，用于控制程序执行逻辑，如程序的分阶段执行以及goroutine的退出。
