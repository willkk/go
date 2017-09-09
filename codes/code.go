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
