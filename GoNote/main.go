package main

import "gonote/GoNote/note"

// noteUtil "gonote/GoNote/util" // 给引入包起别名

// var A = noteUtil.F("note.A")

// func init() {
// 	noteUtil.F("main.init1")
// }
// func init() {
// 	noteUtil.F("main.init2")
// }
// func init() {
// 	noteUtil.F("main.init3")
// }

// 1. 跨包调用
func main() {
	// note.SayHelloWorld()

	// note.VarablesAndConstant()

	// fmt.Println("跨包全局变量")
	// fmt.Println(note.Version)

	// note.Pointer()

	// note.IfElse()
	// note.Swith()

	// note.For()
	// note.LabelAndGoto()
	// note.Function()
	// note.Defer()
	// note.DeferRecord()
	// fmt.Println("note.DeferRecord() 之后还在运行")

	// note.Array()
	// note.Slice()
	// note.Map()
	// note.TypeDefintionAndTypeAlias()
	// var a = note.Test{
	// 	Name: "111",
	// 	Id:   11,
	// }
	// a.Set("kkk")
	// note.Interface()
	// note.Goroutine()
	// note.Channel()
	// note.RandNum()
	// note.StrConv()
	// note.Log()
	// note.PackageSync()
	// note.Recursion(10000)
	// note.Closure()
	note.Sort(100)
}
