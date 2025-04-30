package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 协程定义使用的demo
	routineDemoA()
	testSavePointObject()
}

// 定义安全的对象
type SafePointObject struct {
	// 互斥锁
	me    sync.Mutex
	count int
}

func (s *SafePointObject) increment() {
	s.me.Lock()
	s.me.Unlock()
	s.count++
}

func (s *SafePointObject) getCount() int {
	return s.count
}

func testSavePointObject() {
	var saveObj = SafePointObject{}
	for i := 0; i < 1000; i++ {
		go func() {
			saveObj.increment()
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println(saveObj.getCount())
}

// go-routine轻量级携程，可以启动多个，但是内存变量也是共享的，所以存在内存安全问题，包括import的go文件
func routineDemoA() {
	// 第一个go routine
	go func() {
		fmt.Println("1")
	}()

	go func(str, charStr string) {
		fmt.Println("str : ", str, " charStr : ", charStr)
	}("ni shi yi ge hai zi", " ")

	printNum()
}

func printNum() {
	for i := 0; i < 10; i++ {
		// 睡眠200ms
		time.Sleep(200 * time.Millisecond)
		fmt.Println("current i : ", i)
	}
}
