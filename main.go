// 程序所属包
package main

//导入依赖包
import (
	helloPrint "fmt"
	"time"
	"unsafe"
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

	helloPrint.Println("测试各种类型占用大小")
	var lint int = 1
	println("int", unsafe.Sizeof(lint))
	var lint16 int16 = 1
	println("int16", unsafe.Sizeof(lint16))
	var lint32 int32 = 1
	println("int32", unsafe.Sizeof(lint32))
	var lint64 int64 = 1
	println("int64", unsafe.Sizeof(lint64))
	var luint uint = 1
	println("uint", unsafe.Sizeof(luint))
	var luint8 uint8 = 1
	println("uint8", unsafe.Sizeof(luint8))
	var luint16 uint16 = 1
	println("uint16", unsafe.Sizeof(luint16))
	var luint32 uint32 = 1
	println("uint32", unsafe.Sizeof(luint32))
	var luint64 uint64 = 1
	println("uint64", unsafe.Sizeof(luint64))
	var luintptr uintptr = 1
	println("uintptr", unsafe.Sizeof(luintptr))
	var lcomplex64 complex64 = 1
	println("complex64", unsafe.Sizeof(lcomplex64))
	var lcomplex128 complex128 = 1
	println("complex128", unsafe.Sizeof(lcomplex128))
	var lfloat32 float32 = 1
	println("float32", unsafe.Sizeof(lfloat32))
	var lfloat64 float64 = 1
	println("float64", unsafe.Sizeof(lfloat64))
	var lbyte byte = 1
	println("byte", unsafe.Sizeof(lbyte))
	var lrune rune = 1
	println("rune", unsafe.Sizeof(lrune))
}
