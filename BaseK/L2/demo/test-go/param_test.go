package demo

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println("current dictionary is demo !")
	initStu()

}

type dirT = int

// 定义常量枚举
const (
	FileA dirT = iota
	Dire
	C
)

type Student struct {
	Name     string
	Age      int8
	AddressA string
	//标识这是一个指针存储 ， var dirs = Dirt{File,Dire} typeA = & dirs  *dirs标识typeA参数，可以修改typeA的值
	typeA  *[]dirT
	Iphone []int8
}

func initStu() {
	// 定义的stu 指向对象存放的地址
	var stu = &Student{
		Name: "sss",
		Age:  18,
	}

	//  将值类型传递
	// 指向存放地址的值
	fmt.Println("stu", *stu)

	updateStu(stu)
	// 根据指针变量设置的对象的地址的值，通过*操作地址的值
	fmt.Println(*&stu.Age)
}

func updateStu(stu *Student) {
	fmt.Println("enter update age function ")
	stu.Age = 99
}
