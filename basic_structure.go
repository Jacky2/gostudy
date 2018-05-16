//Go 程序的一般结构: basic_structure.go

// 当前程序的包名
package main

// 导入其他包 使用 . 符号代表 Println 时不需要输入包名。
import . "fmt"

// 常量定义
const PI = 3.14

// 全局变量的声明和赋值
var name = "gopher"

// 一般类型声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由main函数作为程序入口点启动
func main() {
	Println("Hello World!")
}
