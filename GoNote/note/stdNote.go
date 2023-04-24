package note

import (
	"errors"
	"fmt"
	"gonote/GoNote/util"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// 随机数
// 6.1.1 随机数种子 rand.Seed(int64)
// Unix时间: UTC 世界时间标准
// GO获取Unix纳秒 time.Now().UnixNano()
func RandNum() {
	// fmt.Println(rand.Intn(10) - 9)
	// // 获取纳秒
	// fmt.Println("纳秒: ", time.Now().UnixNano())

	// 随机整数
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Intn(10) - 9)
	}
}

// 6.2 字符串类型转换
func StrConv() {
	i1 := 123
	s1 := "测试转换"
	// 其他类型转字符串 用法同fmt.Printf()
	s2 := fmt.Sprintf("%d@%s", i1, s1)
	fmt.Println("其他类型转字符串 : s2= ", s2)

	// 字符串解析维其他类型
	var (
		i2 int
		s3 string
	)
	n, err := fmt.Sscanf(s2, "%d@%s", &i2, &s3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("成功解析了%d个数据\n", n)
	fmt.Println("i2= ", i2)
	fmt.Println("s3= ", s3)

	s4 := strconv.FormatInt(123, 4)
	fmt.Println("s4= ", s4)
	u, err2 := strconv.ParseUint(s4, 4, 0)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("u= ", u)
}

// 字符串操作
func PackageStrings() {
	// 判断是否包含
	fmt.Println(strings.Contains("hello", "o"))

	// 不区分大小写比较字符串
	fmt.Println(strings.EqualFold("hello", "HELLO"))
}

// 中文字符常见操作
func PackageUtf8() {
	fmt.Println(utf8.RuneCountInString("hello,世界"))
}

// 时间常规操作
func PackageTime() {
	fmt.Println("\n6.5.1 时段")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(time.Millisecond * 1000)
	}
	fmt.Println()
	d, err := time.ParseDuration("1000s")
	if err != nil {
		panic(err)
	}
	fmt.Println("d = ", d)
	t, err2 := time.Parse("2006年1月2日 15点4分", "2022年1月1日 18点18分")
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(time.Since(t))
}

// error
func Error() {
	defer func() {
		err := recover()
		fmt.Println("捕捉到了错误 ", err)
	}()
	err := errors.New("可爱的错误")
	fmt.Println("err = ", err)
	err2 := fmt.Errorf("%s的错误", "该死")
	fmt.Println("err2 = ", err2)

	// 自己抛出异常
	panic(err2)
}

// 日志
func Log() {
	defer func() {
		err := recover()
		fmt.Println("捕捉到了错误:", err)
	}()
	err := errors.New("可爱的错误")
	util.INFO.Println(err)
	// util.WARN.Panic(err)
	util.ERR.Fatalln(err)
}

// 单元测试
func IsNotNegative(n int) bool {
	return n > -1
}

// sync 包
func PackageSync() {
	fmt.Println("\n6.14.1 Mutex互斥锁 / 6.14.2 WaitGroup")
	var c int
	var mutex sync.Mutex
	var wg sync.WaitGroup
	primeNum := func(n int) {
		defer wg.Done()
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return
			}
		}
		mutex.Lock()
		c++
		mutex.Unlock()
	}

	for i := 2; i < 100001; i++ {
		wg.Add(1)
		go primeNum(i)
	}
	wg.Wait()
	fmt.Printf("\n共找到%v个素数\n", c)

	fmt.Println("\n6.14.3 Cond 控制多个协程阻塞的能力")
	//  提供了同事控制多个协程阻塞的能力
	cond := sync.NewCond(&mutex)
	for i := 0; i < 10; i++ {
		go func(n int) {
			cond.L.Lock()
			cond.Wait()
			fmt.Printf("协程%v被唤醒了\n", n)
			cond.L.Unlock()
		}(i)
	}

	for i := 0; i < 15; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		if i == 4 {
			fmt.Println()
			cond.Signal()
		}
		if i == 10 {
			fmt.Println()
			cond.Broadcast()
		}
	}

	fmt.Println("\n6.14.4 once 多个协程调用 只会执行一次")
	var once sync.Once
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				fmt.Println("只有一次机会")
			})
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("\n6.14.5 Map 线程安全 ")
	var m sync.Map
	m.Store(1, 100)
	m.Store(2, 200)
	m.Store(3, 300)
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("m[%v] = %v\n", key, value.(int))
		return true
	})
}

// Sort 排序
