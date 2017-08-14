package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"strings"
	"time"
	"net/http"
	"log"
	"image"
)

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(100 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
		default:
			fmt.Println("    .")
			time.Sleep(50*time.Millisecond)
		}
	}
}
func sample_i() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func sample_h() {
	c := make(chan int, 1)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
func sample_g() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func sample_f() {
	go say("hello")
	say("world")
}

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sample_e() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func sample_d() {
	http.Handle("/string", String("I'm a frayed knot."))
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

type String string

func (string String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(string)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func sample_c() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}

type Hello struct {
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func sample_b() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {

	return reader.r.Read(b)
}

func sample_a() {
	reader := strings.NewReader("我Hello, Reader!")
	bytes := make([]byte, 8)
	for {
		n, err := reader.Read(bytes)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, bytes)
		fmt.Printf("b[:n]=%q\n", bytes[:n])
		if err == io.EOF {
			break
		}
	}
}

func sampleZ() {
	err := ErrNegativeSqrt(2).Error()
	fmt.Println(err)
}

type ErrNegativeSqrt float64

func Sqrt2(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := math.Sqrt(x)
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sampleY() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func sampleX() {
	addrs := map[string]IpAddr{
		"lookback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}

type IpAddr [4]byte

func (ipAddr IpAddr) String() string {
	i0 := ipAddr[0]
	i1 := ipAddr[1]
	i2 := ipAddr[2]
	i3 := ipAddr[3]
	return fmt.Sprintf("%v.%v.%v.%v", i0, i1, i2, i3)
}
func sampleW() {
	a := Person{"Lou", 25}
	b := Person{"Bob", 28}
	fmt.Println(a)
	fmt.Println(b)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func sampleV() {
	var w Writer
	w = os.Stdout
	fmt.Fprint(w, "hello world\n")
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReaderWriter interface {
	Reader
	Writer
}

func sampleU() {
	var a Abser
	f := MyFloat(3)
	a = f
	fmt.Println(a.Abs())
	v := Vertex3{2, 3}
	a = &v
	fmt.Println(a.Abs())
	v2 := Vertex3{3, 4}
	a = &v2
	//a = v2 //错误，v2是值类型，不是指针类型
	fmt.Println(v2.Abs())
	fmt.Println(a.Abs())
}

type Abser interface {
	Abs() float64
}

func sampleT() {
	myFloat := MyFloat(2)
	fmt.Println(myFloat.Abs())
	fmt.Println(myFloat)
	myFloat = MyFloat(-2)
	fmt.Println(myFloat.Abs())
	// 对于非指针类型的，是拷贝多一份出来，对之前地址所在的值没有影响
	fmt.Println(myFloat)
	fmt.Println(myFloat.changeMyFloat())
	// 对于指针类型的，可以直接对地址重新赋值
	fmt.Println(myFloat)
}
func sampleS() {
	vertex3 := Vertex3{3, 4}
	v := vertex3
	vertex3.X = 4
	v.Scale(5)
	fmt.Println(v.Abs())
	fmt.Println(vertex3.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		f = MyFloat(32)
		return float64(-f)
	}
	return float64(f)
}

func (f *MyFloat) changeMyFloat() float64 {
	if *f < 0 {
		*f = MyFloat(-13)
		return float64(-*f)
	}
	return float64(*f)
}

type Vertex3 struct {
	X, Y float64
}

func (v *Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func sampleR() {
	// 斐波拉其函数
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		c := a + b
		a = b
		b = c
		return c
	}
}
func sampleQ() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println("i=", i, pos(i), neg(-2*i))
	}
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func sampleP() {
	// 函数也是值
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(2, 8))
}
func sampleO() {
	m := make(map[string]int)
	m["anwser"] = 3
	fmt.Println(m)
	m["anwser"] = 32
	fmt.Println(m)
	delete(m, "anwser")
	fmt.Println(m)
	v, ok := m["anwser"]
	fmt.Printf("The value %v is ok? %v", v, ok)
}
func sampleN() {
	var m map[string]Vertex2
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.3,
		-73.0,
	}
	fmt.Println(m["Bell Labs"])
}

type Vertex2 struct {
	Lat, Long float64
}

func sampleM() {
	var powerArray = []int{1, 2, 4, 8, 16, 32, 63}
	for i := range powerArray {
		fmt.Printf("2**%d = %v\n", i, math.Pow(2, float64(i)))
	}
	for _, v := range powerArray {
		fmt.Println(v)
	}
	_2 := 2
	fmt.Println(_2)
}
func sampleL() {
	var a []int
	a = append(a, 0)
	printSlice("a", a)
	a = append(a, 1)
	printSlice("a", a)
	a = append(a, 1, 2, 3, 4)
	printSlice("a", a)
}
func sampleK() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil")
	}
}
func sampleJ() {
	a := make([]int, 5)
	println(len(a))
	b := make([]int, 0, 6)
	println(len(b))
	println(cap(b))
	b = b[:cap(b)]
	println(len(b))
	printSlice("a", a)
	printSlice("b", b)
}
func printSlice(s string, ints []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(ints), cap(ints), ints)
}
func sampleI() {
	p := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("p====", p)
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
	fmt.Println(p[0:4])
}
func sampleH() {
	p.X = 3
	fmt.Println(v1, v2, v3, p)
}
func sampleG() {
	vertex := Vertex{1, 2}
	x := vertex.X
	y := vertex.Y
	v := &vertex
	z := vertex
	fmt.Println(x, y, vertex.X)
	v.X = 1e9
	// 通过指针的方式可以改变vertex的值
	vertex.X = 2e3
	z.X = 3e2
	// 通过z无法改变vertex值
	fmt.Println(vertex)
	fmt.Println(z)
	fmt.Println(*v)
}

type Vertex struct {
	X int
	Y int
}

func sampleF() {
	i := 3
	p := &i
	println(p)
	println(*p)
}
func sampleE() {
	for i := 0; i < 10; i++ {
		defer println("i=", i)
	}
	println("hi")
}
func sampleD() {
	defer println("hi")
	println("world")
	println("world2")
	println("world3")
}
func sampleC() {
	i := time.Now().Hour()
	switch {
	case i < 7:
		println("hi")
	case i < 8:
		println("hi2")
	case i < 9:
		println("ff")
	default:
		println("default")
	}
}
func f() int {
	return 3
}
func sampleB() {
	switch os := runtime.GOOS; os {
	case "linux":
		println("Linux")
	}
}
func sampleA() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

func Sqrt(x float64) float64 {
	z := 1.0
	z = z - (math.Pow(z, 2)-x)/2*z
	return z
}
func sample9() {
	if v := math.Pow(3, 4); v < 12 {
		fmt.Println("if v=", v)
	} else {
		fmt.Println("else v=", v)
	}
}
func sample8() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("%v", sum)
	for sum < 10000 {
		sum = sum + 10000
	}
	println()
	println(sum)
}
func sample7() {
	fmt.Println(needInt(Small))
	fmt.Println(needInt64(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
func needInt64(x uint64) uint64 {
	return x*10 + 1
}

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float32) float32 {
	return x * 0.1
}

const (
	Big   = 1 << 10
	Small = Big >> 9
)

func sample6() {
	v := 3 + 0.5i
	fmt.Printf("v is of type %T", v)
}
func sample5() {
	var i int = 3
	var f float64 = float64(i)
	fmt.Print(i, f)
}
func sample4() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q", i, f, b, s)
}
func sample3() {
	x, y := split(49)
	fmt.Print(x, y)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sample2() {
	a, b := swap("a", "b")
	fmt.Print(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}

func sample1() {
	fmt.Println("hello, world")
	print(math.Pi)
}
