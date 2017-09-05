
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
