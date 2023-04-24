package util

import "fmt"

var i = 0

// 匿名函数
var F = func(s string) int {
	fmt.Printf("本次被%s调用\n", s)
	i++
	fmt.Printf("工具类被%v次调用\n", i)
	return i
}
