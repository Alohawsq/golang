package test

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("hello test")
}

func Test() {
	fmt.Println("hello world")
	time.Sleep(3 * time.Second)
	fmt.Println("^---^")
}
