package note

import (
	"errors"
	"testing"
)

// x 代表 变量
// go test -v ./.... 即使运行成功也输出日志
// go test -benchtime xhxmxs./... 规定测试时间
// go test -benchtime Nx ./.... 规定测试次数
// go test -cpu N ./.... 制定测试cpu数量
// go test -run Testxxxx ./...制定测试函数

// 单元测试 Test 开头 文件以_test结尾
func TestIsNotNegative(log *testing.T) {
	err := errors.New("Is Negative")
	if IsNotNegative(0) {
		log.Log("OK")
	} else {
		log.Fatal(err)
	}
	if IsNotNegative(1) {
		log.Log("OK")
	} else {
		log.Error(err)
	}
}
