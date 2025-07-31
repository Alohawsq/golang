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

**注意在新增目录后要在根目录下执行``go mod init 目录名称``**，之后会在go.mod文件中新增一个module，此后就可以在文件中使用目录加名称的方式进行引入了。
### import用法
import语句可以用于导入源码文件所依赖的package包；

import不得导入源码文件没有用到的package包；

在编译时会按照import包的顺序进行编译导入，所有包都导入完成后才会对main中的常量和变量进行初始化，然后执行main中的init函数，最后执行main函数；

import语法格式:
```azure
import "fmt"
import "time"

import (
    "fmt"
    "time"
)

```
#### import别名用法
点（.）操作的含义是：点（.）标识导入后，调用该包中函数时可以省略前缀包名称

下划线（\_）操作的含义是：导入该包，但不导入整个包，而执行该包中的init函数，因而无法通过
包名来调用包中的其他函数。使用下划线（\_）操作往往是为了注册包里的引擎，让外部可以方便地使用。
```azure
#  别名用法
import (
    helloPrint "fmt"
)
helloPrint.Println("a")

# 导入点用法
import . "fmt"
Println("aa")

# 下划线用法
import _ "test/test"

会执行test目录中test文件的init函数

```








