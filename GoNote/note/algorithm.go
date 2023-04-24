package note

import (
	"fmt"
	"math/rand"
	"time"
)

// 7.1 递归
// 递归优化 缓存已计算过的数据
var fibonacciRes []int

func fibonacci(n int) int {
	if n < 3 {
		return 1
	}
	if fibonacciRes[n] == 0 {
		fibonacciRes[n] = fibonacci(n-2) + fibonacci(n-1)
	}
	return fibonacciRes[n]
}
func Recursion(n int) {
	fibonacciRes = make([]int, n+1)
	fmt.Printf("第%v位 斐波那契数是%v\n", n, fibonacci(n))
}

// 闭包
func closureFunc() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次调用接收到n=%v\n", n)
		i++
		fmt.Printf("匿名工具函数被第%v次调用\n", i)
		return i
	}
}
func Closure() {
	f := closureFunc()
	f(2)
	f(4)
	f = closureFunc()
	f(6)
}

// 冒泡排序
func bubbleSort(s []int) {
	lastIndex := len(s) - 1
	for i := 0; i < lastIndex; i++ {
		for j := 0; j < lastIndex-i; j++ {
			if s[j] > s[j+1] {
				t := s[j]
				s[j] = s[j+1]
				s[j+1] = t
			}
		}
	}
}

// 选择排序
func selectorSort(s []int) {
	lastIndex := len(s) - 1
	for i := 0; i < lastIndex; i++ {
		min := i
		// max := lastIndex - i
		for j := i + 1; j <= lastIndex; j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		if min != i {
			t := s[i]
			s[i] = s[min]
			s[min] = t
		}
	}
}

// 插入排序
func InsertSort(s []int) {
	for i := 1; i < len(s); i++ {
		t := s[i]
		j := i - 1
		for ; j >= 0 && s[j] > t; j-- {
			s[j+1] = s[j]
		}
		if j != i-1 {
			s[j+1] = t
			// fmt.Println("s=", s)
		}
	}
}

// 快速排序
func quickSort(s []int, leftIndex int, rightIndex int) {
	if leftIndex < rightIndex {
		pivot := s[rightIndex]
		var rs []int
		l := leftIndex
		for i := leftIndex; i < rightIndex; i++ {
			if s[i] > pivot {
				rs = append(rs, s[i])
			} else {
				s[l] = s[i]
				l++
			}
		}
		s[l] = pivot
		copy(s[l+1:], rs)
		if leftIndex < l-1 {
			quickSort(s, leftIndex, l-1)
		}
		if l+1 < rightIndex {
			quickSort(s, l+1, rightIndex)
		}
	}
}

func Sort(n int) {
	s := make([]int, n)
	seedNum := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		rand.Seed(seedNum)
		s[i] = rand.Intn(10001)
		seedNum++
	}
	fmt.Println("排序前:", s)
	// 冒泡排序
	// bubbleSort(s)
	// 选择排序
	// selectorSort(s)
	// InsertSort(s)
	// 快速排序
	quickSort(s, 0, len(s)-1)
	fmt.Println("排序后:", s)
}
