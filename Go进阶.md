# Go进阶

## 1. 编程模式
**a) 面向过程（通过包级别的显隐控制，实现对外接口）：** 将特定目的的计算过程封装成函数，数据以参数形式传入和传出，数据组织能力限制在函数参数。
适用场景：底层工具类模块，不涉及数据的封装，比如数学，加解密，IO等。不适合场景：上层复杂业务逻辑封装，函数太多，不利于业务逻辑和信息流的组织管理。信息组织粒度是函数。

应用举例，encoding/json包定义：  

![方法定义](https://github.com/willkk/go/blob/master/images/goadvance_procdef.png "方法定义")

使用时：  

![工厂方法](https://github.com/willkk/go/blob/master/images/goadvance_procuse.png)  

**b) 面向对象（类型实现基本的成员函数，提供对外的操作接口）：** 将一组逻辑上关联的函数封装成对象，数据作为对象成员被这组函数共享。数据组织能力为类型本身包含的数据块。

适用场景：复杂业务逻辑和流程的封装。不适用场景：底层工具类的封装。信息组织粒度提升到了对象或模块。

应用举例，net/http包Server类型：　
![ServerMux](https://github.com/willkk/go/blob/master/images/goadvance_classservermux.png)  
![Server](https://github.com/willkk/go/blob/master/images/goadvance_classserver.png)  

**使用时，通过包级别的接口对外开放。**

**c) 面向接口：** 将一组逻辑上关联的对象的公共行为抽象出来封装成接口，使用这些接口时需要生成具体的实现类型。数据组织能力为实现该公共行为集合的全部类型。

更高层次的抽象，信息组织粒度提升到模块或整个程序级别。

适用场景：模块级或应用程序及架构设计。不适用场景：底层工具类型，中间层次的对象等小模块类型。

应用举例，hash包：  

![](https://github.com/willkk/go/blob/master/images/goadvance_interfacehash.png)  

hash包里面的adler32，crc32，crc64和fnv分别实现对应接口，构造函数返回对应上述接口类型。  

fnv：  
![](https://github.com/willkk/go/blob/master/images/goadvance_interfacehashfnv.png)  

**使用时，声明interface类型或指针即可指向任意子类型。**

**d) 常见设计模式（主要解决耦合，可扩展性，性能问题）：**

| 设计模式 | 应用举例 | 目的/好处
| --------| -----| ----  
| 单例模式 | json配置文件+config结构体，处理配置文件模块|	中心维护，动态更新，避免重复读取，省去大量配置操作代码。
| 工厂模式 | 使用工厂模式生成系统全部的Commands| 抽象公共Command接口和基类型，定义各子类型；<br />重新组织代码，继承减少重复代码；<br />面向接口编程，提高代码复用度和可扩展性。
| 享元模式 | 使用连接池管理底层mysql和redis数据库连接|降低昂贵资源反复创建的开销，降低端口使用量。

**e) 程序架构（水平拆分+垂直拆分）**

垂直拆分：基于功能，业务，场景等将应用拆分成多个垂直模块。如配置模块，日志模块，消息模块（接入模块，处理模块，响应模块），数据库模块，下游模块。

<table>
<tr><td colspan="6" align="center">应用</td></tr>
<tr height="200"><td>日志模块</td>
<td>配置模块</td>
<td>消息处理模块</td>
<td>数据库模块</td>
<td>平行/下游调用模块</td>
<td>......</td>
</tr>  
</table>

水平拆分：基于流程，先后，上下游等将应用拆分成多个水平层次。如前端的消息接入模块，中间的消息处理模块，后续的消息响应模块，底层的数据库模块，后端的下游模块。

<table>
<tr>
<td rowspan="4" height="300" align="center">应用</td>
<td>消息接入模块</td>  
</tr>
<tr><td>消息路由模块</td></tr>
<tr><td>消息处理模块</td></tr>
<tr><td>数据库访问模块</td></tr>
</table>

#### 消息路由模块  

**net/http：**  

<!--  
flow
st=>start: Server.Serve
op1=>operation: Accept
op2=>operation: go routine
op3=>operation: ServerHandler.ServeHTTP
cond=>condition: User handler is nil?
op4=>operation: ServerMux
op5=>operation: DefaultServeMux.ServeHTTP
op6=>operation: UserHandler.ServeHTTP
op7=>operation: mux.m[req.path].ServeHTTP

e=>end
st->op1->op2->op3->cond
cond(yes)->op5->op7->e
cond(no)->op6->e

-->
![](https://github.com/willkk/go/blob/master/images/goadvance_nethttp2.png)  

**重点是map保存全部的path信息，对进入的request进行过滤，找到对应的handler。如果使用自定义的Handler，用户需要自己完成类似工作。**

**Iris：**

![]()

**不同之处是，**

## 2. 常用第三方库

### a) Redis

分库：避免key冲突，类似c++的namespace。只支持数字分库。

**Sharding方案（分片 http://www.jianshu.com/p/14835303b07e）**

|Sharding方案| 类型|	特点|	优点|	缺点
| -- | -- | -- | -- | -- 
| Redis Cluster | 服务端分片 | 服务端分片：自动计算key的CRC16对16384取模，根据范围分配到特定的节点。 | 1. 客户端代码独立于后端拓扑。<br/>2. 后端拓扑变化对客户端透明。 |1. 需要手动配置和启动多个节点M-S实例节点。<br/>2. 需要使用额外的cluster库。<br/>3. 不使用一致性hash。<br/>4. 总共16384个slot，节点数据量均衡问题。<br/>5. 动态增删节点（节点数据需要从/向其他节点移入/移出）。<br/>6. 不支持多个key的联合操作和事务。
|Redis Sharding|客户端分片|	客户端对key进行hash，确定实际实例节点。	|1. 一致性hash<br/>2. 分片规则灵活。| 1. 增删节点，少量键值迁移。<br/>2. 客户端代码与后端拓扑耦合度高。
|Redis Proxy|代理分片	|代理层进行hash分片，实现客户端和服务端的无感。	|1. 客户端和服务端的无感，伸缩性好。<br/>2. 可共享后端redis连接。|1. 连接性能有一定损失。

**高可用方案（Sentinel）：**

Sentinel使用分布式方式监控和管理Redis实例，包括：监控，通知，自动故障切换，服务发现（返回具体master）。Sentinel服务的基本单元是单个实例的主从结构，即1个Master+多个Slaves，充当客户端与Redis间的服务中介。可以将Sentinel看做Redis网络的Wrapper，对外提供一个高可用的分布式网络架构。单个实例的主从之间数据是重复的。Sentinel部署条件：需要至少3个Sentinel实例；分开部署；客户端库支持。典型拓扑(目标是保证足够的Sentinel数量)：　

### b) Mysql

驱动包：github.com/go-sql-driver/mysql

orm包：gorm/xorm

ORM：对象关系映射，就是把数据库表映射成一个类型（C++的类，Go的struct等），对底层数据库进行逻辑封装，方便数据库操作。

https://beego.me/docs/mvc/model/object.md

定义：
type User struct {
    Id          int
    Name        string
    Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
    Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
    Id          int
    Age         int16
    User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
    Id    int
    Title string
    User  *User  `orm:"rel(fk)"`    //设置一对多关系
    Tags  []*Tag `orm:"rel(m2m)"`
}
CRUD：
C：
o := orm.NewOrm()
var user User
user.Name = "slene"
user.IsActive = true

id, err := o.Insert(&user)
if err == nil {
    fmt.Println(id)
}
R：
o := orm.NewOrm()
user := User{Id: 1}

err := o.Read(&user)
U：
o := orm.NewOrm()
user := User{Id: 1}
if o.Read(&user) != nil {
    user.Name = "MyName"
    if num, err := o.Update(&user); err == nil {
        fmt.Println(num)
    }
}
D：
o := orm.NewOrm()
if num, err := o.Delete(&User{Id: 1}); err == nil {
    fmt.Println(num)
}
### c) RabbitMQ(作用, 优缺点)

异步，可靠。

　　　　
### d) RESTful/RPC（http://www.ruanyifeng.com/blog/2014/05/restful_api.html）

|方案 |实现|	优点|	缺点
| -- | -- | -- |--
|RESTful(REpresentational State Transfer)| 通过特定格式的URL标识资源，资源的操作抽象为多个状态的转移。（统一表示+状态转移）|1. 统一格式标识资源。<br/>2. 常用的CRUD操作。<br/>3. 适合公共服务资源调用。|1. 性能低。<br/>2. 仅适合可进行状态抽象的业务。 
|RPC(Remote Procedure Call)|通过特定的消息协议格式进行远程过程调用。（预定函数名称和参数是+远程过程调用）|1. 消息格式可以是文本或二进制，原文或经过压缩。<br/>2. 性能高。<br/>3. 适合轻量型数据操作，即RPC的本意，函数+参数。<br/>4. 适合企业内部服务间通信。|1. 双方需要提前商定消息协议格式。 

**RESTful典型应用（api.github.com）：**

**API认证服务：**

**RPC典型应用：**

　　　　
