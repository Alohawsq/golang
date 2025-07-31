// 程序所属包
package main

//导入依赖包
import (
	helloPrint "fmt"
	"time"
)

// 常量定义
const NAME = "hello"

// 变量定义
var a = "world"

// 一般类型声明
type helloInt int

// 结构体声明
type Learn struct {
	name string
	age  int
}

// 接口声明
type Ilearn interface {
}

// 函数定义
func learnHello() {
	helloPrint.Println("lean hello")
}

// main()函数
func main() {
	learnHello()
	helloPrint.Println(a)
	helloPrint.Println(NAME)
	var learn = Learn{
		name: "Lili",
		age:  18,
	}
	helloPrint.Println(learn.name, learn.age)
	helloPrint.Println(time.Now())
}
