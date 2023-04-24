package note

import (
	"fmt"
	"sync"
	// "gonote/GoNote/util"
)

var v = 100

// 跨包调用全局常量 首字母大写)
const (
	Version int = 1000
)

// var A = util.F("note.A")

// func init() {
// 	util.F("note.init1")
// }
// func init() {
// 	util.F("note.init2")
// }
// func init() {
// 	util.F("note.init3")
// }

// 2.1 转义字符
func SayHelloWorld() {
	fmt.Println("\"hello world!\"")
}

// 2.2 变量与常量 (都可以设置全局)
func VarablesAndConstant() {
	// 类型推断
	var v1 int // 整数类型默认值 0
	var v2 int = 2
	var v3 = 3
	v1 = 1
	v4 := 4
	var (
		v5     = 5
		v6 int = 6
		v7 int
	)
	fmt.Printf("v1 = %v, v2= %v, v3=%v, v4= %v, v5= %v, v6= %v,v7= %v\n", v1, v2, v3, v4, v5, v6, v7)

	fmt.Println("\n 常量")
	const (
		c1 = 8
		c2 = iota // 当前行数 (注意 行数是从0开始计算)
		c3 = iota
		c4 // 默认值
		c5 = 12
		c6
	)
	fmt.Printf("c1 = %v, c2= %v, c3=%v, c4= %v, c5= %v, c6= %v\n", c1, c2, c3, c4, c5, c6)

	fmt.Println("\n 全局变量")
	v = 200
	fmt.Println(v)
}

// 2.3
func dataType() {

}

// 2.4 指针
// * (取值内存地址符号)
// & (获取值内存地址符号)
func increase(n int) {
	n++
	fmt.Printf("increase 结束时 n = %v", n)
}
func increase2(n *int) {
	*n++
	fmt.Printf("increase 结束时 n = %v", n)
}

func Pointer() {
	var src = 2023

	increase(src)
	// 获取内存地址 可以保证
	increase2(&src)
}

// 3.1 if..else
func IfElse() {
	var age uint8
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)
	if age < 13 {
		fmt.Println("小朋友不要学编程")
	}
}

// 3.2 swith
func Swith() {
	var age uint8
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)
	switch {
	case age < 13:
		fmt.Println("小朋友不要学编程")
		fallthrough
	case age < 10:
		fmt.Println("小小朋友不要学编程")
	default:
		fmt.Println("大朋友不要学编程")
	}
}

// 3.3
func For() {
	fmt.Println("\n 3.3 无限循环")
	i := 1
	for {
		fmt.Print(i, "\t")
		i++
		if i == 11 {
			fmt.Println()
			break
		}
	}

	fmt.Println("\n 3.3 条件循环")
	i = 1
	for i < 11 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println()

	fmt.Println("\n 3.3 标准for循环")
	for i := 1; i < 11; i++ {
		fmt.Print(i, "\t")
	}
	fmt.Println()
}

// 3.4 label/goto 标签
func LabelAndGoto() {
outside:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("+ ")
			if i == 9 && j == 4 {
				break outside
			}
		}
		fmt.Println()
	}
}

// 3.5 函数
// 1. 参数: 函数可以接受0～多个参数， 形参类型相同时，可以简写不确定形参个数时可以用…数据类型来生成形参切片
// 2. 返回: golang函数支持0～多个返回值 可以命名返回值（会被视作函数顶部定义的变量
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func getSumAndDiffrence(n1, n2 int) (int, int) {
	sum := n1 + n2
	diffrence := n2 - n1
	return sum, diffrence
}

// 简写 不用定义返回值
func getSumAndDiffrence2(n1, n2 int) (sum, diffrence int) {
	sum = n1 + n2
	diffrence = n2 - n1
	return
}

func Function() {
	fmt.Println("单一返回值")
	res := getSum(2, 3)
	fmt.Println("res = ", res)

	fmt.Println("多返回值")
	res1, res2 := getSumAndDiffrence(2, 3)
	fmt.Println("结果 sum =  ", res1, ", diffrence = ", res2)

	fmt.Println("多返回值 (简写)")
	res1, res2 = getSumAndDiffrence2(2, 3)
	fmt.Println("结果 sum =  ", res1, ", diffrence = ", res2)

	// 匿名函数
	var getRes = func(n1, n2 int) (sum, diffrence int) {
		sum = n1 + n2
		diffrence = n2 - n1
		return
	}
	res1, res2 = getRes(2, 3)
	fmt.Println("结果 sum =  ", res1, ", diffrence = ", res2)
	fmt.Printf("getRes=%v,Type of getRes=%T\n", getRes, getRes)
}

// 3.6 延迟执行 defer
// 下面这种命名方式 闭包
func deferUtil() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次调研接收到n=%v\n", n)
		i++
		fmt.Printf("匿名工具函数被%v次调用\n", i)
		return i
	}
}

func Defer() int {
	f := deferUtil()
	// 延迟执行的函数会被压入栈中, return 后按照先进后出的顺序调用
	// 延迟执行的函数其参数会立即求值
	defer f(f(3))
	return f(2)
}

// defer recover
func DeferRecord() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}

// 3.7 init函数
// 1. 每个包都可以有自己的init函数, 且可以有多个
// 2. 执行顺序(取决于包的依赖关系):
// 被依赖包的全局变量 ----> 被依赖包的init函数 ----> .... -----> main包的全局变量 ----> main的init函数 -----> main的函数

// 3.8 包的

// 3.9 数组
func Array() {
	// 4.1.1 声明 长度不能为空
	var a [3]int = [...]int{1, 456, 789}
	a[0] = 123
	// for 遍历
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%v]=%v\n", i, a[i])
	}

	// for range 遍历 _ 这个可以代替
	for i, v := range a {
		fmt.Printf("a[%v]=%v\n", i, v)
	}

	// 多维数组

}

// 切片 是对数组的使用
func Slice() {
	array := [5]int{1, 2, 3, 4, 5}
	var s1 []int = array[1:4] // [开始引用的下标:结束引用的下标+1] 如果从第0项开始 可以不用输入、如果引用结果 结尾可以不用输入
	fmt.Println("array=", array)
	s1[0] = 0
	fmt.Println("array=", array)
	s2 := s1[1:]
	s2[0] = 0
	fmt.Println("array=", array)
	var s3 []int
	fmt.Println("s3==nil?", s3 == nil)

	// make 直接分配内存空间
	// make([]type, len, cap) cap默认与len相同
	s3 = make([]int, 3, 5)
	fmt.Printf("len(s3) = %v, cap(s3)=%v ", len(s3), cap(s3))

	// 更简单的声明方式 自动创建底层数组
	s4 := []int{1, 2, 3}
	fmt.Println("s4=", s4)

	fmt.Println("\n4.2.3 动态追加元素")
	// 底层将创建了新的数组,不再引用原数组
	s1 = append(s1, 6, 7, 8)
	s1[4] = 0
	fmt.Println("array=", array)
	fmt.Println("s1=", s1)

	// 添加其他切片,底层创建一个新的切片
	s5 := append(s1, s2...)
	fmt.Println("s5=", s5)

	s6 := []int{1, 1}
	copy(s5, s6)
	// 容量能接收多少就接收多少
	fmt.Println("数组复制: s5=", s5)
}

// 4.3 map
func Map() {
	var m1 map[string]string
	fmt.Println("m1 == nil ?", m1 == nil)
	// make(Type, 初始size)
	m1 = make(map[string]string, 2)
	m1["早上"] = "敲代码"
	m1["中午"] = "送外卖"
	m1["晚上"] = "开滴滴"
	fmt.Println("m1 =", m1)

	m2 := map[string]string{
		"下午": "改BUG",
		"凌晨": "卖早餐",
	}
	fmt.Println("m2 =", m2)

	v, ok := m2["早上"]
	if ok {
		fmt.Println("v =", v)
	} else {
		fmt.Println("key不存在")
	}

	// 删除元素
	delete(m1, "晚上")
	fmt.Println("m1 =", m1)
	// 重新滞空
	m2 = make(map[string]string)
	fmt.Println("m1 =", m2)

	for k, v := range m1 {
		fmt.Printf("m1[%v]=%v", k, v)
	}
}

// 自定义数据类型
func TypeDefintionAndTypeAlias() {
	fmt.Println("\n 4.4.1 自定义数据")
	type mesType uint16
	var u1000 uint16 = 1000
	// 混用需要类型转换
	var textMes mesType = mesType(u1000)
	fmt.Printf("textMes=%v, type = %T", textMes, textMes)
	fmt.Println("\n 4.4.1 自定义数据别名")
}

// 结构体
type User struct {
	Name string
	Id   uint32
}
type Account struct {
	User
	password string
}
type Contact struct {
	*User
	Remark string
}

func Struct() {
	var u1 User = User{
		Name: "张三",
	}
	u1.Id = 1000
	var u2 *User = new(User)
	u2.Id = 10001

	// var a1 = Account{
	// 	User:     u1,
	// 	password: "66666",
	// }
	// var c1 *Contact = &Contact{
	// 	User:   u2,
	// 	Remark: "王麻子",
	// }
}

type Ty interface{ Struct() string }
type Test struct {
	Name string
	Id   uint32
}

func (name *Test) Set(name1 string) {
	name.Name = name1
	fmt.Printf("name.Name" + name1)
}

// func (*Test) Struct() string {

// 	fmt.Println("111")
// 	return "111"
// }

var c1, c2, c3 chan int
var i1, i2 int

// select语句
func Select() {
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	}
}

// 接口
type Msg interface {
	setType()
}

type textMsg struct {
	Type string
	Text string
}

type imgMsg struct {
	Type string
	Text string
}

func (tm *textMsg) setText() {
	tm.Text = "开始表演了"
}
func (im *imgMsg) setText() {
	im.Text = "月下遛鸟图"
}
func (tm *textMsg) setType() {
	tm.Type = "文字消息"
}
func (im *imgMsg) setType() {
	im.Type = "图片消息"
}
func SendMsg(m Msg) {
	m.setType()
	// 还原始类型
	switch mprt := m.(type) {
	case *textMsg:
		mprt.setText()
	case *imgMsg:
		mprt.setText()
	}
	fmt.Println("m=", m)
}

// 可以定义空接口 可以保存任何类型
func Interface() {
	tm := textMsg{}
	SendMsg(&tm)

	im := imgMsg{}
	SendMsg(&im)

	var n1 int = 1
	n1interface := interface{}(n1)
	// 类型断言 返回
	n2, ok := n1interface.(int)
	if ok {
		fmt.Println("n2=", n2)
	} else {
		fmt.Println("类型断言失败")
	}
}

// nil 问题
// nil值 有类型没有值, 接口本身不是nil 可以进行处理
// nil接口 既没有保存值,也灭有保存类型, 使用就会报错

// 5.3 协程
// Goroutine 是Go运行管理的轻量级线程
// 主线程结束时,协程会被中断,需要有效的阻塞机制

// 多线程同时对一个内存空间进行写操作会导致数据竞争, 可以使用sync包下 互斥锁Mutex(不常用) go可以使用channel来解决这个问题

// 定义互斥锁
var (
	c    int
	lock sync.Mutex
)

func PrimeNum(n int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	fmt.Printf("%v\t", n)
	lock.Lock()
	c++
	lock.Unlock()
}
func Goroutine() {
	for i := 2; i < 100001; i++ {
		go PrimeNum(i)
	}
	var key string
	fmt.Scanln(&key)
	fmt.Printf("\n共找到%v个素数\n", c)
}

// 5.4 channel 管道
// 5.4.1 声明与存取
// 定义类型 make(Type, (缓冲区容量))
// 不带缓冲区的管道必须结合协程使用
// 可以查看长度len(channel) 或容量cap(channel)
// 存入 channel <- value
// 取出 value <- channel
//丢弃 <- channel
// 原则上先进先出 自动阻塞
// 数据要保持流动, 否则会阻塞报错

// 5.4.2 关闭 可以通过close(channel)

// 5.4.3 for 或 for...range 不断从channel接收值, 知道它被关闭(缺乏关闭机制会报错)

// 5.4.4 select...case 适用于无法确认合适关闭信道的情况
// 通常结合for循环使用
// 会阻塞到某一个分支可以继续执行时执行该分支,当没有可以执行的会执行default 分支
func pushNum(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	// 不关闭 一定会报错 deadlock
	close(c)
}
func Channel() {
	// 如果不放缓冲区 (10) make(chan int, 10) 往管道里面添加数据 会报死锁 are asleep - deadlock
	// var c1 chan int = make(chan int)
	// 使用协程调用
	// go pushNum(c1)
	// for v := range c1 {
	// 	fmt.Printf("%v\t", v)
	// }
	// 定义带缓冲channel
	var c2 chan int = make(chan int, 100)
	for i := 0; i < 100001; i++ {
		go pushPrimeNum(i, c2)
	}
Print:
	for {
		select {
		case v := <-c2:
			fmt.Printf("%v\t", v)
		default:
			fmt.Println("所有素数已经找到")
			break Print
		}
	}
}

func pushPrimeNum(n int, c chan int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	c <- n
}
