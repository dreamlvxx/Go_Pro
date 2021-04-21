package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strings"
	"text/template"
	"time"
)

type Circle int32
type Cube int32

var myVar int32
var myBo bool = true

var (
	aa    int32
	bb    bool
	com64 complex64
)

const (
	MyCircle Circle = 1
	MyCube   Cube   = 2
)
const GoUsage = `Go is a tool for managing Go source code.

Usage:
    go command [arguments]`

const sad = `asdadasd[]
sadasd
sad
`

func main() {
	fmt.Println("start connect mysql")
	//db,_:=sql.Open("mysql","root:root@(127.0.0.1:3306)/golang")
	//
	//err :=db.Ping()
	//if err != nil{
	//	fmt.Println("数据库链接失败")
	//}

	fmt.Println("start connect mysql")

	fmt.Println("-------------------------------------")
	var b int
	b = 7
	fmt.Println(b)
	fmt.Println("hello world")
	fmt.Println("   sss")
	var bo bool
	fmt.Println(bo)
	var er error
	fmt.Println(er)
	var s = true
	fmt.Println(s)

	myV1 := ":=var1"
	fmt.Println(myV1)
	vart, varr := "asd,", "asd"
	fmt.Println(vart + varr)
	_, qwe := "nouse", "qwe"
	fmt.Println(qwe)

	const a = "consta"

	const mycom = 1 + 2i
	mycom1 := 1 + 2i

	fmt.Println(mycom1)

	const (
		Unknow = 0
		Male   = 1
		FeMale = 2
	)
	fmt.Println(Unknow)
	q := 6
	j, k := add(b, q)

	var balance = [5]int{1000, 2, 3, 17, 50}
	hharr(balance[:])
	fmt.Println(j, k)

	pointVal := 6
	var pointer *int
	pointer = &pointVal
	fmt.Printf("地址是 %x", pointer)
	fmt.Printf("     地址的值是 %x", *pointer)
	var ptr *int
	fmt.Println(ptr == nil)

	var stud Student

	stud.id = 123
	stud.name = "xxx"
	stud.sex = "boy"

	var stupointer *Student
	stupointer = &stud

	fmt.Println(stupointer)
	fmt.Println(stud.sex)
	changeStudMesg(stupointer)
	fmt.Println(stud.sex)

	var slice = []int{1, 2, 3}
	var sum int
	for _, num := range slice {
		sum += num
	}
	fmt.Println(sum)

	mymap := map[string][]string{"111": {"1233,123"}}
	mi := []string{"asd", "aaaa"}
	mymap["asd"] = mi
	fmt.Println(mymap["111"])
	delete(mymap, "111")

	for k, v := range mymap {
		fmt.Println(k, "--", v)
	}

	fmt.Println(os.Args[2])
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go pNums(url, ch) // start a goroutine
	}
	fmt.Println("start+++++++++++++++++++++++++++++++++++++++++++++")
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Println("end+++++++++++++++++++++++++++++++++++++++++++++++")

	fmt.Println(sad)

	sliceInt := []int{1, 2, 3, 4}
	sliceInt = append(sliceInt[:], 1, 2)
	reversSlice(sliceInt)
	fmt.Println(sliceInt)
	//myServer()
	resTemplete()
	myFuncVar = myFunc
	myFuncVar(5)
	ga := Galaxy{}
	ga.ExtraMethod(2)

	gg := &Galaxy{}
	gg.ExtraFunc()
	fmt.Println(gg.X, gg.Y)

	var pptr *Galaxy
	pptr = gg
	pptr.ExtraFunc()
	fmt.Println(pptr.Y, pptr.X)
	//
	//ra := new(Rocket)
	//time.AfterFunc(10 * time.Second, func() { ra.Launch(2) })
	//
	//rocket := Rocket{}
	//receiver := Rocket.Launch
	//receiver(rocket,2)
	//AfterFunc(2,ra.Launch)//传入方法值
	//myintr := MyInterStruct{}
	//MyInterMethod(myintr)
}

func MyInterMethod(inter MyInter) {

}

type MyInterStruct struct {
}

func (m MyInterStruct) Sleep(a int) int {
	return 0
}

type MyInter interface {
	Sleep(a int) int
}

func AfterFunc(d int, f func(a int)) {

}

type Rocket struct { /* ... */
}

func (r *Rocket) Launch(a int) { /* ... */ }

type Galaxy struct {
	X, Y int
}

func (g Galaxy) ExtraMethod(a int) int {
	fmt.Println("this is a extra method")
	return 0
}

var myFuncVar func(int) (int, int)

func myFunc(a int) (int, int) {
	return 0, 0
}

func (g *Galaxy) ExtraFunc() {
	g.X = 4
	g.Y = 5
}

type Humen struct {
	X int32
	Y int32
	s int
}

type Point struct {
}

type People struct {
	Humen
	S     int32
	point Point
}

func (h Humen) getHumenX() {
	h.Y = 4
	h.s = 6
}

func getNUmm() {
	poe := People{Humen: Humen{1, 2, 2}, S: 22, point: Point{}}
	fmt.Println(poe.X)
}

func reversSlice(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func changeStudMesg(stu *Student) {
	stu.sex = "girl"
}

type Student struct {
	name string
	sex  string
	id   int32
}

func add(a, b int) (int, int) {
	var s = a + b
	fmt.Println(s)
	return b, a
}

func hharr(arr []int) {
	runtime.Gosched()
}

func readFromFile(name string) {
	file, _ := os.Open(name)
	input := bufio.NewScanner(file)

	for input.Scan() {
		fmt.Println(input.Text())
	}
}

func readFromFile1(name string) {
	data, _ := ioutil.ReadFile(name)
	for _, line := range strings.Split(string(data), "\n") {
		fmt.Println(line)
	}
}

func rawReadFromfile(file *os.File) {
	var sad []byte
	file.Read(sad)
}

func TimeNow() {
	rand.Seed(time.Now().UTC().UnixNano())
	//anim := gif.GIF{LoopCount: 12}
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func getResFromNet(url string) {
	res, _ := http.Get(url)
	if res != nil {
		data, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		fmt.Printf("%s", data)
	} else {
		os.Exit(1)
	}

}

func pNums(mes string, ch chan<- string) {
	ch <- mes + "puffix"
}

func myServer() {
	http.HandleFunc("/111/", HandleJsonFromClient)
	http.ListenAndServe("localhost:8000", nil)
}

func HandlerFromClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func HandleJsonFromClient(w http.ResponseWriter, r *http.Request) {
	items := []Issue{{Id: 1, Description: "1111"}, {Id: 2, Description: "2222"}}
	res := IssueResult{
		Item: &items,
	}
	json.NewEncoder(w).Encode(res)
}

type IssueResult struct {
	Item *[]Issue
}
type Issue struct {
	Id          int32
	Description string
}

type TempleteSelf struct {
	Count int32
	Items *[]Issue
}

const templ = `{{.Count}} issues:
{{range .Items}}----------------------------------------
Id: {{.Id}}
Des:   {{.Description}}
{{end}}`

//模板
func resTemplete() {
	var report = template.Must(template.New("issuelist").
		Parse(templ))
	items := []Issue{{Id: 1, Description: "1111"}, {Id: 2, Description: "2222"}}
	res := TempleteSelf{
		Count: 12,
		Items: &items,
	}
	if err := report.Execute(os.Stdout, res); err != nil {
		log.Fatal(err)
	}
}
