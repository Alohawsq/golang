// 包名称 package 只能出现在文件开头
package main

// 导入包
import . "fmt"

// 常量定义
const NAME_A = "hello"

// 变量定义
var aa = "world"

// main函数
func main() {
	var b string = "code"
	Println(b)
	Println(NAME_A)
	Println(aa)
}
