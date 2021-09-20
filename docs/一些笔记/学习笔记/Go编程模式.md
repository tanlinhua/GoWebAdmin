# GO编程模式

## 1. 基础编码

### Slice的底层实现(Slice Internal)

关于Slice的实现，我之前有[一讲](https://github.com/Junedayday/code_reading/blob/master/basic/data_struct.md#slice)专门分析过底层实现。考虑到很多朋友没有细看，那我就再简单地讲一下。

```go
type slice struct {
	array unsafe.Pointer // Slice底层保存数据的指针
	len int // 当前使用的长度
	cap int // 分配的长度
}
```

掌握Slice的底层实现，能让你真正理解一些看似“奇怪的”现象：

```go
func main(){
    foo := make([]int, 5)
    foo[3] = 42
    foo[4] = 100
    
    bar := foo[1:4]
    bar[1] = 99
    
    fmt.Println(foo)
    // [0 0 99 42 100]
    fmt.Println(bar)
    // [0 99 42]
}
```

> Tip: bar和foo是共享slice结构体底层的array的，所以修改了bar数组，foo也会变化

```go
func main(){
    a := make([]int, 32)
    b := a[1:16]
    
    a = append(a, 1)
    a[2] = 42
  
    fmt.Println(b)
    // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}
```

> Tip: a和b原来是共享array的，但在a = append(a, 1)后发生了扩容，a和b指向的array发生了变化

```go
 func main(){
    path := []byte("AAAA/BBBBBBBBB")
    sepIndex := bytes.IndexByte(path,'/')
    dir1 := path[:sepIndex]
    dir2 := path[sepIndex+1:]
    fmt.Println(cap(dir1),cap(dir2))
    // 14 9
    fmt.Println("dir1 =>",string(dir1))
    // dir1 => AAAA
    fmt.Println("dir2 =>",string(dir2))
    // dir2 => BBBBBBBBB
    
    dir1 = append(dir1,"suffix"...)
    fmt.Println("dir1 =>",string(dir1))
    // dir1 => AAAAsuffix
    fmt.Println("dir2 =>",string(dir2))
    // dir2 => uffixBBBB
}
```

> Tip: 核心点在于理解dir1和dir2的cap分别是14和9。由于dir1的当前len=4，append的长度=6，4+6<14，所以不会发生扩容



### 深度对比(Deep Comparison)

我们先看一下示例，`data`结构体中四个注释为`not comparable`表示无法直接用 == 符号对比

```go
type data struct {
	num    int               // ok
	checks [10]func() bool   // not comparable
	doit   func() bool       // not comparable
	m      map[string]string // not comparable
	bytes  []byte            // not comparable
}

func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2))
	// prints: v1 == v2: true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))
	// prints: m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
	// prints: s1 == s2: true
}
```

> Tip： 示例比较复杂，其实要表达的内容比较简单：
>
> 函数、map、切片（不包括数组）以及它们的复合结构（如函数的数组），无法直接对比，只能用 reflect.DeepEqual



### 函数传参VS对象方法(Function vs Receiver)

```go
type Person struct {
	Name   string
	Sexual string
	Age    int
}

func PrintPerson(p *Person) { fmt.Printf("Name=%s, Sexual=%s, Age=%d\n", p.Name, p.Sexual, p.Age) }
func (p *Person) Print()    { fmt.Printf("Name=%s, Sexual=%s, Age=%d\n", p.Name, p.Sexual, p.Age) }

func main() {
	var p = Person{
		Name: "Hao Chen", Sexual: "Male", Age: 44,
	}

	PrintPerson(&p)
	// Name=Hao Chen, Sexual=Male, Age=44
	p.Print()
	// Name=Hao Chen, Sexual=Male, Age=44
}
```

> Tip: 示例比较简单，但其中蕴含的意义非常大，如对Person这个对象的抽象、简化代码等。
>
> 另外值得一提的是，Go编译器会根据方法 `func (p *Person) Print()` 的定义，将 `p.Print()`中的p从`Person`转换为`*Person`。



### 面向接口编程(Interface Patterns)

这个模块非常重要，希望大家倒一杯水，细细品尝。

示例是一个很简单的interface实现，用来打印接口，我们看看代码。

```go
type Country struct {
	Name string
}

type City struct {
	Name string
}

type Printable interface {
	PrintStr()
}

func (c Country) PrintStr() {
	fmt.Println(c.Name)
}

func (c City) PrintStr() {
	fmt.Println(c.Name)
}

func main() {
	c1 := Country{"China"}
	c2 := City{"Beijing"}

	var cList = []Printable{c1, c2}
	for _, v := range cList {
		v.PrintStr()
	}
}
```

那么，这时问题来了，如果我要实现N个`Printable`，就要定义N个strcut+N个`PrintStr()`方法。

前者的工作不能避免，而后者能否简化？那么示例来了

```go
type WithName struct {
	Name string
}

type Country struct {
	WithName
}

type City struct {
	WithName
}

type Printable interface {
	PrintStr()
}

func (c WithName) PrintStr() {
	fmt.Println(c.Name)
}

func main() {
	c1 := Country{WithName{"China"}}
	c2 := City{WithName{"Beijing"}}

	var cList = []Printable{c1, c2}
	for _, v := range cList {
		v.PrintStr()
	}
}
```

> Tip: 核心就是用 embedded 的特性来删除冗余的代码。当然，代价是初始化会稍微麻烦点。

---

这时候，陈皓又给出了一个例子，即打印的内容会根据具体的实现不同时，无法直接用`WithName`来实现

```go
type Country struct {
	Name string
}

type City struct {
	Name string
}

type Printable interface {
	PrintStr()
}

func (c Country) PrintStr() {
	fmt.Println("Country:", c.Name)
}

func (c City) PrintStr() {
	fmt.Println("City:", c.Name)
}

func main() {
	c1 := Country{"China"}
	c2 := City{"Beijing"}

	var cList = []Printable{c1, c2}
	for _, v := range cList {
		v.PrintStr()
	}
}
```

首先，我们要明确是否有必要优化。如果只有示例中这么几行代码，我们完全没必要改写。那如果系统真复杂到一定程度，我们该怎么办呢？

这是一个很发散性的问题，我这里给出一个个人比较常用的解决方案，作为参考。

```go
type WithTypeName struct {
	Type string
	Name string
}

type Country struct {
	WithTypeName
}

func NewCountry(name string) Printable {
	return Country{WithTypeName{"Country", name}}
}

type City struct {
	WithTypeName
}

func NewCity(name string) Printable {
	return City{WithTypeName{"City", name}}
}

type Printable interface {
	PrintStr()
}

func (c WithTypeName) PrintStr() {
	fmt.Printf("%s:%s\n", c.Type, c.Name)
}

func main() {
	c1 := NewCountry("China")
	c2 := NewCity("Beijing")

	var cList = []Printable{c1, c2}
	for _, v := range cList {
		v.PrintStr()
	}
}
```

> Tip： 这种方法的好处有很多（先不谈弊端），比如可以将具体的实现`Country`和`City`私有化，不对外暴露实现细节。今天不做细谈。


最后，送上一句经典：

**Program to an interface, not an implementation**

### 时间格式(Time)

这部分的内容实战项目中用得不多，大家记住耗子叔总结出来的一个原则即可：

> 尽量用`time.Time`和`time.Duration`，如果必须用string，尽量用`time.RFC3339`



然而现实情况并没有那么理想，实际项目中用得最频繁，还是自定义的`2006-01-02 15:04:05`

```go
time.Now().Format("2006-01-02 15:04:05")
```



### 性能1(Performance1)

#### Itoa性能高于Sprint

主要性能差异是由于`Sprint`针对的是复杂的字符串拼接，底层有个buffer，会在它的基础上进行一些字符串的拼接；

而`Itoa`直接通过一些位操作组合出字符串。

```go
// 170 ns/op
func Benchmark_Sprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(rand.Int())
	}
}

// 81.9 ns/op
func Benchmark_Itoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(rand.Int())
	}
}
```



#### 减少string到byte的转换

主要了解go的`string`到`[]byte`的转换还是比较耗性能的，但大部分情况下无法避免这种转换。

我们注意一种场景即可：从`[]byte`转换为`string`，再转换为`[]byte`。

```go
// 43.9 ns/op
func Benchmark_String2Bytes(b *testing.B) {
	data := "Hello world"
	w := ioutil.Discard
	for i := 0; i < b.N; i++ {
		w.Write([]byte(data))
	}
}

// 3.06 ns/op
func Benchmark_Bytes(b *testing.B) {
	data := []byte("Hello world")
	w := ioutil.Discard
	for i := 0; i < b.N; i++ {
		w.Write(data)
	}
}
```



#### 切片能声明cap的，尽量初始化时声明

了解slice的扩容机制就能很容易地理解。切片越长，影响越大。

```go
var size = 1000

// 4494 ns/op
func Benchmark_NoCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make([]int, 0)
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
}

// 2086 ns/op
func Benchmark_Cap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make([]int, 0, size)
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
}
```



#### 避免用string做大量字符串的拼接

频繁拼接字符串的场景并不多，了解即可。

```go
var strLen = 10000

// 0.0107 ns/op
func Benchmark_StringAdd(b *testing.B) {
	var str string
	for n := 0; n < strLen; n++ {
		str += "x"
	}
}

// 0.000154 ns/op
func Benchmark_StringBuilder(b *testing.B) {
	var builder strings.Builder
	for n := 0; n < strLen; n++ {
		builder.WriteString("x")
	}
}

// 0.000118 ns/op
func Benchmark_BytesBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for n := 0; n < strLen; n++ {
		buffer.WriteString("x")
	}
}
```



### 性能2(Performance2)

#### 并行操作用sync.WaitGroup控制

#### 热点内存分配用sync.Pool

注意一下，一定要是`热点`，千万不要 **过早优化**


#### 倾向于使用lock-free的atomic包

除了常用的`CAS`操作，还有`atomic.Value`的`Store`和`Load`操作，这里简单地放个实例：

```go
func main() {
	v := atomic.Value{}
	type demo struct {
		a int
		b string
	}

	v.Store(&demo{
		a: 1,
		b: "hello",
	})

	data, ok := v.Load().(*demo)
	fmt.Println(data, ok)
	// &{1 hello} true
}
```

复杂场景下，还是建议用`mutex`。



#### 对磁盘的大量读写用bufio包

`bufio.NewReader()`和`bufio.NewWriter()`



#### 对正则表达式不要重复compile

```go
// 如果匹配的格式不会变化，全局只初始化一次即可
var compiled = regexp.MustCompile(`^[a-z]+[0-9]+$`)

func main() {
	fmt.Println(compiled.MatchString("test123"))
	fmt.Println(compiled.MatchString("test1234"))
}
```



#### 用protobuf替换json

go项目内部通信尽量用`protobuf`，但如果是对外提供api，比如web前端，`json`格式更方便。



#### map的key尽量用int来代替string

```go
var size = 1000000

// 0.0442 ns/op
func Benchmark_MapInt(b *testing.B) {
	var m = make(map[int]struct{})
	for i := 0; i < size; i++ {
		m[i] = struct{}{}
	}
	b.ResetTimer()
	for n := 0; n < size; n++ {
		_, _ = m[n]
	}
}

// 0.180 ns/op
func Benchmark_MapString(b *testing.B) {
	var m = make(map[string]struct{})
	for i := 0; i < size; i++ {
		m[strconv.Itoa(i)] = struct{}{}
	}
	b.ResetTimer()
	for n := 0; n < size; n++ {
		_, _ = m[strconv.Itoa(n)]
	}
}
```

示例中`strconv.Itoa`函数对性能多少有点影响，但可以看到`string`和`int`的差距是在数量级的。



### Further

PPT中给出了8个扩展阅读，大家根据情况自行阅读。

如果说你的时间只够读一个材料的话，我推荐大家反复品读一下[Effective Go](https://golang.org/doc/effective_go.html)

## 2. 继承与嵌入

### 嵌入和委托(Embedded)

接口定义

```go
// 定义了两种interface
type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}
```

Label 实现了 Painter

```go
// 标准组件，用于嵌入
type Widget struct {
	X, Y int
}

// Label 实现了 Painter
type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

func (label Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}
```

ListBox实现了Painter和Clicker

```go
// ListBox声明了Paint和Click，所以实现了Painter和Clicker
type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}

func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}

func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}
```

Button也实现了Painter和Clicker

```go
// Button 继承了Label，所以直接实现了Painter
// 接下来，Button又声明了Paint和Click，所以实现了Painter和Clicker，其中Paint方法被覆
type Button struct {
	Label // Embedding (delegation)
}

func (button Button) Paint() { // Override
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}

func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}
```

方法调用

```go
func main() {
	label := Label{Widget{10, 70}, "Label"}
	button1 := Button{Label{Widget{10, 70}, "OK"}}
	button2 := Button{Label{Widget{50, 70}, "Cancel"}}
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	for _, widget := range []interface{}{label, listBox, button1, button2} {
		// 默认都实现了Painter接口，可以直接调用
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
	}
}
```

这个例子代码很多，我个人认为重点可以归纳为一句话：

**用嵌入实现方法的继承，减少代码的冗余度**



耗子叔的例子很精彩，不过我个人不太喜欢`interface`这个数据类型（main函数中），有没有什么优化的空间呢？

```go
// 定义两种方法的组合
type PaintClicker interface {
	Painter
	Clicker
}

func main() {
	// 在上面的例子中，interface传参其实不太优雅，有没有更优雅的实现呢？那就用组合的interface
	for _, widget := range []PaintClicker{listBox, button1, button2} {
		widget.Paint()
		widget.Click()
	}
}
```



### 反转控制(IoC)

先看一个Int集合的最基本实现

```go
// Int集合，用于最基础的增删查
type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(x int) { set.data[x] = true }

func (set *IntSet) Delete(x int) { delete(set.data, x) }

func (set *IntSet) Contains(x int) bool { return set.data[x] }
```

现在，需求来了，我们希望对这个Int集合的操作是可撤销的

```go
// 可撤销的Int集合，依赖于IntSet，我们看看基本实现
type UndoableIntSet struct { // Poor style
	IntSet    // Embedding (delegation)
	functions []func()
}

func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

// 新增
// 不存在元素时：添加元素，并新增撤销函数：删除
// 存在元素时：不做任何操作，并新增撤销函数：空
func (set *UndoableIntSet) Add(x int) { // Override
	if !set.Contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() { set.Delete(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

// 删除，与新增相反
// 存在元素时：删除元素，并新增撤销函数：新增
// 不存在元素时：不做任何操作，并新增撤销函数：空
func (set *UndoableIntSet) Delete(x int) { // Override
	if set.Contains(x) {
		delete(set.data, x)
		set.functions = append(set.functions, func() { set.Add(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

// 撤销：执行最后一个撤销函数function
func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil // For garbage collection
	}
	set.functions = set.functions[:index]
	return nil
}
```

上面的实现是一种顺序逻辑的思路，整体还是挺麻烦的。有没有优化思路呢？

定义一下Undo这个结构。

```go
type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}
```

细品一下这里的实现：

```go
type IntSet2 struct {
	data map[int]bool
	undo Undo
}

func NewIntSet2() IntSet2 {
	return IntSet2{data: make(map[int]bool)}
}

func (set *IntSet2) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet2) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSet2) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet2) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}
```

我们看一下，这块代码的前后逻辑有了啥变化：

1. 之前，撤销函数是在Add/Delete时添加的，函数中包含了IntSet的操作，也就是 **Undo依赖IntSet**
2. 而修改之后，撤销函数被抽象为Undo，撤销相关的工作直接调用Undo相关的工作即可，也就是 **IntSet依赖Undo**

我们再来分析一下

- Undo是控制逻辑 - 撤销动作
- IntSet是业务逻辑 - 保存数据的功能。

**业务逻辑依赖控制逻辑，才能保证在复杂业务逻辑变化场景下，代码更健壮！**

## 3. 错误处理

### 函数式处理(Functional)

```go
type Number struct {
	a int
	b string
	c bool
	d []int32
	e error
}

func (n *Number) parse(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, &n.a); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &n.b); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &n.c); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &n.d); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &n.e); err != nil {
		return err
	}
	return nil
}
```



引入了函数式编程的方式，我们看看有什么改变

```go
func (n *Number) parse(r io.Reader) error {
	// 先定义一个error
	var err error

	// 定义函数，注意这里的err的作用域是来自上面定义的
	read := func(data interface{}) {
		// 先检查error，如果已经有错误则不检查
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	// 注意，这个函数的调用逻辑和之前的差别在于一点：
	// 即使前面的发生了error，下面的函数也会被调用
	read(&n.a)
	read(&n.b)
	read(&n.c)
	read(&n.d)
	read(&n.e)

	return err
}
```



### 对象嵌入错误(ErrorObject)

先看一个标准库中的实现

```go
func main() {
	input := bytes.NewReader([]byte("hello"))
	
	// 扫描数据，这里不会直接返回错误
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		token := scanner.Text()
		fmt.Println(token)
	}
	
	// 从Err()方法中获取错误
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
```

它的根本思想，是将`error`嵌入到了对象中。那我们借鉴一下

```go
type Reader struct {
	r   io.Reader
	err error
}


func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

func (n *Number) parse(reader io.Reader) error {
	r := Reader{r: reader}

	r.read(&n.a)
	r.read(&n.b)
	r.read(&n.c)
	r.read(&n.d)
	r.read(&n.e)

	return r.err
}
```

捎带提一句：个人不太喜欢上面`scanner`的错误处理方式，这个要求使用方对这个包很熟悉，否则很容易忘掉后面的错误处理逻辑。但后面处理错误的逻辑，就很直接地将错误返回，可读性很强。

### 错误包装(Wrap)

耗子叔给的例子是调用了`github.com/pkg/errors`下的wrap包，不过我更倾向于直接用原生的。

```go
func main() {
	// 原始 error
	err := errors.New("level 1")
	fmt.Println(err)
	// level 1

	// wrap一下error，注意error的占位符是%w
	wraped := fmt.Errorf("%v: %w", "level 2", err)
	fmt.Println(wraped)
	// level 2: level 1

	// unwrap 后获得原来的错误
	unwraped := errors.Unwrap(wraped)
	fmt.Println(unwraped)
	// level 1

	// 过度unwrap会导致错误变成nil
	unwraped2 := errors.Unwrap(unwraped)
	fmt.Println(unwraped2)
	// nil
}
```

但在实际项目实践中，`Wrap`的这个特性并不好用：

如何Wrap Error，在多人协同开发、多模块开发过程中，很难统一。而一旦不统一，容易出现示例中的过度Unwrap的情况。

所以，我认为与其花大精力在制定错误的标准上，还不如利用`fmt.Errorf`将错误信息直观地表述出来。

## 4. 函数式选项

### 一个常见的HTTP服务器(ServerConfig)

我们先来看看一个常见的HTTP服务器的配置，它区分了2个必填参数与4个非必填参数

```go
type ServerCfg struct {
	Addr     string        // 必填
	Port     int           // 必填
	Protocol string        // 非必填
	Timeout  time.Duration // 非必填
	MaxConns int           // 非必填
	TLS      *tls.Config   // 非必填
}

// 我们要实现非常多种方法，来支持各种非必填的情况，示例如下
func NewServer(addr string, port int) (*Server, error)                                   {}
func NewTLSServer(addr string, port int, tls *tls.Config) (*Server, error)               {}
func NewServerWithTimeout(addr string, port int, timeout time.Duration) (*Server, error) {}
func NewTLSServerWithMaxConnAndTimeout(addr string, port int, maxconns int, timeout time.Duration, tls *tls.Config) (*Server, error) {}
```



### 拆分可选配置(SplitConfig)

编程的一大重点，就是要 `分离变化点和不变点`。这里，我们可以将必填项认为是不变点，而非必填则是变化点。

我们将非必填的选项拆分出来。

```go
type Config struct {
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type Server struct {
	Addr string
	Port int
	Conf *Config
}

func NewServer(addr string, port int, conf *Config) (*Server, error) {
	return &Server{
		Addr: addr,
		Port: port,
		Conf: conf,
	}, nil
}

func main() {
	srv1, _ := NewServer("localhost", 9000, nil)

	conf := Config{Protocol: "tcp", Timeout: 60 * time.Second}
	srv2, _ := NewServer("localhost", 9000, &conf)

	fmt.Println(srv1, srv2)
}
```

到这里，其实已经满足大部分的开发需求了。那么，我们将进入今天的重点。



### 函数式选项(Functional Option)

```go
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

// 定义一个Option类型的函数，它操作了Server这个对象
type Option func(*Server)

// 下面是对四个可选参数的配置函数
func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

// 用到了不定参数的特性，将任意个option应用到Server上
func NewServer(addr string, port int, options ...Option) (*Server, error) {
	// 先填写默认值
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	// 应用任意个option
	for _, option := range options {
		option(&srv)
	}
	return &srv, nil
}

func main() {
	s1, _ := NewServer("localhost", 1024)
	s2, _ := NewServer("localhost", 2048, Protocol("udp"))
	s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))

	fmt.Println(s1, s2, s3)
}
```

耗子哥给出了6个点，但我感受最深的是以下两点：

1. 可读性强，将配置都转化成了对应的函数项option
2. 扩展性好，新增参数只需要增加一个对应的方法

那么对应的代价呢？就是需要编写多个Option函数，代码量会有所增加。



如果大家对这个感兴趣，可以去看一下Rob Pike的[这篇blog](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html) 。



### 更进一步(Further)

顺着耗子叔的例子，我们再思考一下，如果配置的过程中有参数限制，那么我们该怎么办呢？

首先，我们改造一下函数Option

```go
// 返回错误
type OptionWithError func(*Server) error
```

然后，我们改造一下其中两个函数作为示例

```go
func Protocol(p string) OptionWithError {
	return func(s *Server) error {
		if p == "" {
			return errors.New("empty protocol")
		}
		s.Protocol = p
		return nil
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) error {
		if timeout.Seconds() < 1 {
			return errors.New("time out should not less than 1s")
		}
		s.Timeout = timeout
		return nil
	}
}
```

我们再做一次改造

```go
func NewServer(addr string, port int, options ...OptionWithError) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	// 增加了一个参数验证的步骤
	for _, option := range options {
		if err := option(&srv); err != nil {
			return nil, err
		}
	}
	return &srv, nil
}
```

改造基本到此完成，希望能给大家带来一定的帮助。

## 5. 映射,递归,过滤

### 映射、规约与过滤(Map/Reduce/Filter)

```go
func MapUpCase(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapLen(arr []string, fn func(s string) int) []int {
	var newArray = []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}

func Filter(arr []string, fn func(n string) bool) []string {
	var newArray = []string{}
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}

func main() {
	var list = []string{"Hao", "Chen", "MegaEase"}

	// 元素一对一映射 string->string
	x := MapUpCase(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)
	// [HAO CHEN MEGAEASE]

	// 元素一对一映射 string->int
	y := MapLen(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", y)
	// [3 4 8]

	// 归约：多个元素->一个元素
	z := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", z)
	// 15

	// 过滤：过滤不满足条件的元素
	f := Filter(list, func(s string) bool {
		return len(s) > 3
	})
	fmt.Printf("%v\n", f)
	// [Chen MegaEase]
}

```



### 应用场景探索(Scenarios)

- `Map` 是一对一的场景，是 **循环中对数据加工处理** 
- `Reduce` 是多对一，是 **数据聚合处理**
- `Filter`是过滤的处理，是 **数据有效性**



我们以常见的账单统计相关的功能，我们会遇上大量的此类情况：

1. 统计消费总额 - Reduce
2. 统计用户A - Filter
3. 统计本月 - Filter
4. 费用转化为美金 - Map

在综合各个因素后，就是大量复杂的、管道式的Map/Reduce/Filter操作。

> 延伸思考一下，这块和SQL语句非常类似



### 泛型(Generic)

耗子叔在接下来的部分，展示了用`reflect`处理泛型情况。我这边简单地截取Map部分解析一下

```go
func TransformInPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

// map的转换函数，slice为切片，function为对应的函数，inPlace表示是否原地处理
func transform(slice, function interface{}, inPlace bool) interface{} {
	// 类型判断，必须为切片
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	// 函数的签名判断，即函数的入参必须和slice里的元素一致
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("trasform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	// 如果是原地，则直接处理函数，结果会保存到入参中（这时入参一般为指针）
	// 如果非原地，那就需要新建一个切片，用来保存结果
	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}
	return sliceOutType.Interface()

}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	// 类型判断
	if fn.Kind() != reflect.Func {
		return false
	}

	// 入参数量和函数签名一致，出参必须只有一个
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}

	// 每个函数入参的类型校验
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}

	// 出参类型的校验
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}
```

仔细阅读这一块代码，我们能学到很多反射方面的知识，尤其是并不常用的函数相关的。

但是，我不建议大家在实际项目中直接使用这一块代码，毕竟其中大量的反射操作是比较耗时的，尤其是在延迟非常敏感的web服务器中。

如果我们多花点时间、直接编写指定类型的代码，那么就能在编译期发现错误，运行时也可以跳过反射的耗时。

## 6. 代码生成

### 简单脚本准备(Simple Script)

为了让大家快速了解这块，我们从一个最简单的例子入手。

#### 模板文件(Template)

首先创建一个模板Go文件，即容器模板：container.tmp.go

```go
package PACKAGE_NAME
type GENERIC_NAMEContainer struct {
    s []GENERIC_TYPE
}
func NewGENERIC_NAMEContainer() *GENERIC_NAMEContainer {
    return &GENERIC_NAMEContainer{s: []GENERIC_TYPE{}}
}
func (c *GENERIC_NAMEContainer) Put(val GENERIC_TYPE) {
    c.s = append(c.s, val)
}
func (c *GENERIC_NAMEContainer) Get() GENERIC_TYPE {
    r := c.s[0]
    c.s = c.s[1:]
    return r
}
```

#### 运行脚本(Shell)

生成的shell脚本，gen.sh

```shell
#!/bin/bash

set -e

SRC_FILE=${1}
PACKAGE=${2}
TYPE=${3}
DES=${4}
#uppcase the first char
PREFIX="$(tr '[:lower:]' '[:upper:]' <<< ${TYPE:0:1})${TYPE:1}"

DES_FILE=$(echo ${TYPE}| tr '[:upper:]' '[:lower:]')_${DES}.go

sed 's/PACKAGE_NAME/'"${PACKAGE}"'/g' ${SRC_FILE} | \
    sed 's/GENERIC_TYPE/'"${TYPE}"'/g' | \
    sed 's/GENERIC_NAME/'"${PREFIX}"'/g' > ${DES_FILE}
```

四个参数分别为

- 源文件名
- 包名
- 类型
- 文件后缀名

#### 入口文件(Generate File)

最后，增加一个创建代码的go文件。

```go
//go:generate ./gen.sh ./template/container.tmp.go gen uint32 container
func generateUint32Example() {
    var u uint32 = 42
    c := NewUint32Container()
    c.Put(u)
    v := c.Get()
    fmt.Printf("generateExample: %d (%T)\n", v, v)
}

//go:generate ./gen.sh ./template/container.tmp.go gen string container
func generateStringExample() {
    var s string = "Hello"
    c := NewStringContainer()
    c.Put(s)
    v := c.Get()
    fmt.Printf("generateExample: %s (%T)\n", v, v)
}
```

#### 运行原理(Generation)

我们运行一下 `go generate`，就能产生对应的文件。

1. 运行go generate，工具会扫描所有的文件
2. 如果发现注释有带 go:generate的，会自动运行后面的命令
3. 通过命令生成的代码，会在源文件添加提示，告诉他人这是自动生成的代码，不要编辑

因此，我们不仅仅可以用`shell`脚本，也可以用各种二进制工具来生成代码。值得一提的是，像Kubernetes这种重量级的项目，大量地应用了这种特性。后面我也会和大家分享在开发web项目中的应用。



下面，我也来介绍几个个人认为比较有用的工具。



### 类型替换工具genny

源项目链接：https://github.com/cheekybits/genny

#### Go文件示例

```go
package queue

import "github.com/cheekybits/genny/generic"

// NOTE: this is how easy it is to define a generic type
type Something generic.Type

// SomethingQueue is a queue of Somethings.
type SomethingQueue struct {
  items []Something
}

func NewSomethingQueue() *SomethingQueue {
  return &SomethingQueue{items: make([]Something, 0)}
}
func (q *SomethingQueue) Push(item Something) {
  q.items = append(q.items, item)
}
func (q *SomethingQueue) Pop() Something {
  item := q.items[0]
  q.items = q.items[1:]
  return item
}
```

#### 脚本

```shell
cat source.go | genny gen "Something=string"
```

官方示例还是采用的是shell脚本，建议替换到 go:generate 中，这样的代码更统一

#### 原理

可以简单地理解成一个类型替换的工具（PS：擅长用sed脚本的朋友也可直接通过shell脚本实现）



### 任意文件转Go(go-bindata)

源网站链接：https://github.com/go-bindata/go-bindata

go-bindata的功能是将任意格式的源文件，转化为Go代码，使我们无需再去打开文件读取了。

这个工具多用在静态网页转化为Go代码（不符合前后端分离的实践），所以具体的使用方式我就不细讲了，大家有兴趣的可以自行阅读教程。

但它有两个优点值得我们关注：无需再进行文件读取操作、压缩。



### 字符串生成工具stringer

stringer是官方提供一个字符串工具，我个人非常推荐大家使用

文档链接：https://pkg.go.dev/golang.org/x/tools/cmd/stringer 

#### Go文件

```go
package painkiller

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
```

#### 脚本

```go
//go:generate stringer -type=Pill
```

于是，就会生成对应的方法`func (Pill) String() string`，也就是直接转化成了其命名。

#### 价值

Go语言在调用 `fmt` 等相关包时，如果要将某个变量转化为字符串，默认会寻找它的`String()`方法。

这时，**良好的命名** 能体现出其价值。尤其是在错误码的处理上，无需再去查询错误码对应的错误内容，直接可以通过命名了解。


## 7. 装饰,管道,访问者模式

### 装饰模式(Decoration)

代码实例

```go
func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Started")
		f(s)
		fmt.Println("Done")
	}
}
```

一句话解释：**在函数f前后，添加装饰性的功能函数，但不改变函数本身的行为**。

这种设计模式，对一些被高频率调用的代码非常有用：

1. HTTP Server被调用的handler
2. HTTP Client发送请求
3. 对MySQL的操作

而装饰性的功能，常见的有：

1. 打印相关的日志信息（Debug中非常有用！）
2. 耗时相关的计算
3. 监控埋点



### 管道模式(Pipeline)

代码示例

```go
type HttpHandlerDecorator func(http.HandlerFunc) http.HandlerFunc
func Handler(h http.HandlerFunc, decors ...HttpHandlerDecorator) http.HandlerFunc {
    for i := range decors {
        d := decors[len(decors)-1-i] // iterate in reverse
        h = d(h)
    }
    return h
}
```

一句话解释：**用不定参数的特性，将入参中的函数，逐个应用到对象上**

> 看到这里，如果你能想起之前 `Functional Option` 那篇，会发现有这块的影子。

主要应用于： 有多种可选择的配置（对应Field）或处理（对应方法）的复杂对象。

耗子叔在后面又增加了一些用Goroutine+Channel的方式，其实就是讲Channel作为一个管道的承载体。



### 访问者模式(Visitor)

关于访问者设计者模式，我之前在Kubernetes源码分析中专门分析了源码。今天，我们也简单地过一下。

```go
// 定义访问的函数类型
type VisitorFunc func(*Info, error) error

// Visitor接口设计
type Visitor interface {
	Visit(VisitorFunc) error
}

// 资源对象
type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

// 将Visitor函数应用到资源对象上
func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}
```

然后看其中一个实现：NameVisitor，其余的也类似，这样就能注入对应的Visitor

```go
type NameVisitor struct {
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		// 这里是运行入参中的VisitorFunc，这一块的逻辑有点像pipeline
		err = fn(info, err)
		// NameVisitor自己实现的Visit逻辑
		return err
	})
}
```

当然，Kubernetes中的Visitor还有进一步的封装，包括遇到错误时的处理，这里不细讲，有兴趣的朋友可以看看我对那一篇的分析。

Visitor模式最大的优点就是 `解耦了数据和程序`。回头看Kubernetes的Visitor应用场景，主要是从各种输入源中解析出资源`Info`。这个过程中Info是数据，各类解析方法是资源。

所以，我认为Visitor模式比较适合的是：**目标数据明确，但获取数据的方法多样且复杂**。但由于多层Visitor调用复杂，建议大家可以在外面再简单地封一层，提供常用的几种Visitor组合后的接口，供使用方调用。

### [更多设计模式](https://github.com/senghoo/golang-design-pattern)