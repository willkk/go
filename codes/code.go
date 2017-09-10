// go spec1 OMIT
// 右侧是简省形式
var tel string = "110" // tel := "110"
var man Man = Man{Name:"Tom", Age:20} // man := Man{Name:"Tom", Age:20}
var man_ages map[string]int8 = map[string]int8{
	"Tom": 20,
	"Will": 18,
}
// 多个变量同时声明，放在括号里
var (
	age int8
	name string
	man1 Man
)
// go spec2 OMIT

// go struct1 OMIT
// 基本类型成员
type Man struct {
	// public
	Name string   `json:"name"`   // 姓名
	Age  int8     `json:"age"`	  // 年龄
	// private
	addr string   `json:"addr"`   // 住址
	// 错误，递归声明
	//Father Man    `json:"father"` // 父亲
}
// 结构体成员
type Student struct {
	Man Man				   // 继承
	// 重复定义
	//*Man
	Teachers []Man	   // 教师
	Scores   [10]int8  // 分数
}
//类型别名，identical to Man
type Human = Man
type MyTime = time.Time
// go struct2 OMIT

// go func1 OMIT
// 空参数，空返回值
func Func() {
}
// 可变参数
func Func2(args ...int) int {
	return 0
}
// 匿名参数及返回值
func Func3( int, string) (map[int]string, string){
	return nil, ""
}
// 命名参数和返回值,可直接使用返回值列表里面的命名参数,并返回空即可。
func Func4(a int, b string) (kvs map[int]string, msg string){
	kvs = nil
	msg = "hello"
	return
}
// 返回值是函数
func Func5() func(a, b int) int {
	return Add
}
// go func2 OMIT


// go for1 OMIT
// for循环，三种等效
for ; ;  {
	time.Sleep(time.Second)
}
for true {
	time.Sleep(time.Second)
}
for  {
	time.Sleep(time.Second)
}
var urls []string = []string{"", "", "", ""}
for url := range urls {
	fmt.Println(url)
}
var chUrl chan string = make(chan string, 10)
// 如果生产者调用close(chUrl), for循环可以继续读取，最后退出
for url := range chUrl {
	fmt.Println(url)
}
// go for2 OMIT


// go defer3 OMIT
// 普通代码
func foo(key string) ([]byte, error) {
    var resp interface{}
    if err1 != nil { 
        json, err := json.Marshal(resp)
        if err != nil {
            return nil, err
        }
        return json, nil
    }
    if err2 != nil { 
        json, err := json.Marshal(resp)
        if err != nil {
            return nil, err
        }
        return json, nil
    }
    json, err := json.Marshal(resp)
    //if err != nil {  ... }
    return json, nil
}

// go defer4 OMIT


// go defer5 OMIT
func foo(key string) (json []byte, err error) {
    var resp interface{}
    defer func() {
        json, err = json.Marshal(resp)
    }()

    if err1 != nil {
        // resp = xx
        return
    }
    // ...
    if err2 != nil {
        // resp = yy
        return
    }

    // resp = zz
    return
}

// go defer6 OMIT


// go defer1 OMIT
// defer后面的函数会在当前作用域退出之前被执行
lock  := new(sync.Mutex)
{
	lock.Lock()
	defer lock.Unlock()
	// codes
	// ......
}

func TestDefer() (ret int){
	defer func() {
		ret = 12
	}()

	ret = 20
	return
}
// go defer2 OMIT

// go no named var1 OMIT
// 普通代码
func foo(key string) (string, int, map[int]interface{}, func(string, int) error) {
    var (
        str string
        integer int
        maps map[int]interface{}
        func1 func (string, int) error
    )

    if err1 != nil {
        //str, integer, maps, func1 = xx, xx, xx, xx 
        return str, integer, maps, func1
    }
    // ...
    if err2 != nil {
        //str, integer, maps, func1 = yy, yy, yy, yy 
        return str, integer, maps, func1
    }

    return str, integer, maps, func1
}

// go no named var2 OMIT

// go named var1 OMIT
func foo(key string) (name string, age int, friends map[string]interface{}, 
                      bar func(string, int) error) {
    // Init return values

    if err1 != nil {
        //name, age, friends, bar = xx, xx, xx, xx
        return
    }
    // ...
    if err2 != nil {
        //name, age, friends, bar = yy, yy, yy, yy
        return
    }

    return
}
// go named var2 OMIT


// go switch OMIT
// 空表达式，默认bool类型true
switch  {
case false: fmt.Println("false")
case true: fmt.Println("true")
case len("dd") > 0: fmt.Println("you can not see me.")
}
x, y := 3, 10
// 默认case自带break, fallthrough取消break
switch  {
case x > y: fmt.Println("x > y")
case x < y: fmt.Println("x < y"); fallthrough
// 多个条件用逗号间隔
case x + 2 < y, x*2 < y : fmt.Println("x+2 < y")
default: fmt.Println("all false")
}
// 检查变量数据类型
switch x.(type) {
case nil:
case chan string:
case func(int, string) string:
}
// go switch2 OMIT

// go select1 OMIT
var ch1, ch2, ch3 chan int
select {
// 多个case满足时，随机选择一个;如果一个也没有，选择default；否则阻塞
case ch1 <- 12:
case ret := <- ch2:
	fmt.Println(ret)
// more为false表示channel关闭
case ret, more := <- ch3:
	fmt.Println(ret, more)
// 超时
case time.After(time.Second*30):
default:
	fmt.Println("get nothing")
}

select {
// 阻塞
}
// go select2 OMIT


// go camel1 OMIT
const (
	blockSize = 512
	TypeSymlink       = '2'    // symbolic link
)

type Header struct {
	ChangeTime time.Time // status change time
	Xattrs     map[string]string
}

const (
	fileNameSize       = 100 // Maximum number of bytes in a standard tar name.
	fileNamePrefixSize = 155 // Maximum number of ustar extension bytes.
)

type headerFileInfo struct {
	h *Header
}

func (fi headerFileInfo) Size() int64        { return fi.h.Size }
func (fi headerFileInfo) ModTime() time.Time { return fi.h.ModTime }
// go camel2 OMIT



// go jsonconf1 OMIT
type ServiceConfig struct {
	Ip string    	`json:"ip"`			// 服务监听地址
	Port string     `json:"port"`		// 服务监听端口
	// ...
}

type MysqlConfig struct {
	Addr string    `json:"addr"`	// mysql服务地址
}

type LogConfig struct {
	File string    `json:"file"`	// log配置文件路径
}

type ServerConfig struct {
	Service ServiceConfig    `json:"service"`// 服务配置
	Mysql   MysqlConfig      `json:"mysql"`  // mysql配置
	Redis 	RedisConfig      `json:"redis"`  // Redis配置
	Log   	LogConfig        `json:"log"`	 // log配置
	// ...
}
// go jsonconf2 OMIT


// go jsonfile1 OMIT
{
  "service": {
    "ip": "0.0.0.0",
    "port": "8888",
    "max_procs": 8
  },
  "mysql": {
    "addr": "10.11.10.12:3306",
    "timeout": 60
  },
  "redis": {
    "addr": "127.0.0.1:6379"
  },
  "log": {
    "file": "../conf/log.conf",
    "timeout": 30
  }
}
// go jsonfile2 OMIT

// go jsonparse1 OMIT

appConfig := flag.String("appConfig", "../config/server.config",
	"")
file, err := os.Open(*appConfig)
if err != nil {
	return
}
data, err := ioutil.ReadAll(file)
if err != nil {
	return
}
srvConfig := &ServerConfig{}
err = json.Unmarshal(data, srvConfig)

max_procs := srvConfig.Service.MaxProcs

// go jsonparse2 OMIT


// Redis1 OMIT
typedef struct dict {
    dictType *type;
    void *privdata;
    dictht ht[2]; //双buffer  
    int rehashidx;
    int iterators;
} dict;
// Redis2 OMIT

func Compose(f, g func(x float) float)
                  func(x float) float {
     return func(x float) float {
        return f(g(x))
    }
}

print(Compose(sin, cos)(0.5))

// Interface1 OMIT
type Reader interface {
    Read([]byte) (int, error)
}
// Interface2 OMIT
// ByteReader implements an io.Reader that emits a stream of its byte value.
type ByteReader byte

func (b ByteReader) Read(buf []byte) (int, error) {
    for i := range buf {
        buf[i] = byte(b)
    }
    return len(buf), nil
}
// Interface3 OMIT
type LogReader struct {
    io.Reader
}

func (r LogReader) Read(b []byte) (int, error) {
    n, err := r.Reader.Read(b)
    log.Printf("read %d bytes, error: %v", n, err)
    return n, err
}
// Interface4 OMIT
var r io.Reader = ByteReader('A')
r = io.LimitReader(r, 1e6)
r = LogReader{r}
io.Copy(ioutil.Discard, r)
// Interface5 OMIT
io.Copy(ioutil.Discard, LogReader{io.LimitReader(ByteReader('A'), 1e6)})
//Interface6 OMIT

// HandlerFunc1 OMIT
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
// HandlerFunc2 OMIT


// Model1 OMIT
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
// Model2 OMIT

o := orm.NewOrm()
var user User
user.Name = "slene"
user.IsActive = true

id, err := o.Insert(&user)
if err == nil {
    fmt.Println(id)
}
// Model3 OMIT
o := orm.NewOrm()
user := User{Id: 1}

err := o.Read(&user)
// Model4 OMIT
o := orm.NewOrm()
user := User{Id: 1}
if o.Read(&user) != nil {
    user.Name = "MyName"
    if num, err := o.Update(&user); err == nil {
        fmt.Println(num)
    }
}
// Model5 OMIT
o := orm.NewOrm()
if num, err := o.Delete(&User{Id: 1}); err == nil {
    fmt.Println(num)
}
// Model6 OMIT
