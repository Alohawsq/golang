# golang
go语言入门学习

## 常用命令
```azure
go build: 用于编译源码文件、代码包、依赖包；

go run: 可以编译并运行Go源码文件；

go install:用来动态获取远程代码包；
```

## 常见IDE
Sublime text2、liteide、Goland

## 基础语法
### 关键字和标识符
25个常用关键字：

| break    | default     | func   | interface | select |
|----------|-------------|--------|-----------|--------|
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

36个预定的标识符：

| append    | bool       | type   | cap   | close   | complex |
|----------|------------|--------|-------|---------|---------|
| complex64   | complex128 | unit16 | copy  | false   | float32 |
| float64     | imag       | int    | int8  | int16   | unit32  |
| int32     | int64      | iota   | len   | make    | new     |
| nil    | panic      | unit64 | print | println | real    |
| recover | string     | TRUE   | unit  | unit8   | unitprt |

### 注释
```azure
// 这是单行注释

/*
这是多行注释
 */
```
### 基础结构
```azure
// 包名称 package 只能出现在文件开头
package main

// 导入包
import "fmt"

// 常量定义
const NAME = "hello"

// 变量定义
var a = "world"

// main函数
func main() {
	var b string = "code"
	fmt.Println(b)
	fmt.Println(NAME)
	fmt.Println(a)
}

```

### package
package是最基本的分发单位和工程管理中依赖关系的体现；

每个Go语言源代码文件开头都必须要有一个package声明，表示源码文件所属代码包；

要生成Go语言可执行程序，必须要有main的package包，且必须在该包下有main函数；

同一路径下只能存在一个package，一个package可以拆成多个源文件组成；
