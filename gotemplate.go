package main

import (
	"fmt"
)

const c = "C" //隐式类型定义
const initTips = "...init..."

var v int = 5 //显式类型定义

type T struct {
}

func init() {
	fmt.Println(c)
	fmt.Println(initTips)
}

func Func1(a, b int) int {
	fmt.Println("Func1")
	//exported function Func1
	//...
	return a + b
}

func main() {
	fmt.Println("main")
	var a int
	a = v
	sumResult := Func1(a, v)
	fmt.Sprintf("sum:%s", sumResult)
	//...
	fmt.Println(a)
}

func (t T)Method1() {
	//...
}

/*
Go 程序的执行（程序启动）顺序如下：

按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：
如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。
*/

