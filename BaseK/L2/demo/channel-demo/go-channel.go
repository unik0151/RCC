package main

import (
	"fmt"
	"time"
)

// channel 一种类型，专门用于goroutine之间通信的线程安全的数据结构
// 在 goroutine中向一个channel 发送数据，另一个goroutine中接受数据
// channel 要满足先进先出原则

/**
- 定义方式 ：
1. var channel1 chan type type可以是结构体也可以是基础数据,int uint float rune string,bool,byte,mapp[type]type, []type
2. channel1 := make(chane type)
3. channel1 := make(chane type,3)
- channel三种操作，发送数据，接受数据，关闭通道
1.  将数据输入通道： channel_name <- var_value
2. 接受数据：value , ok := <- channel_name
 value := <- channel_name

 3. 关闭通道
 close(channel_name)


 锁和channel
 当需要channel协作时,常用channel,而不是Mutex或RWMute的互斥锁
 大部分：流程是根据数据驱动的，channel被使用更频繁
 channel擅长数据流动的场景
 1. 传递数据到其他协程
 2.分发任务，每个任务都是一个数据
 3. 交流异步结果，结果是一个数据
 锁使用：偏向于同一时间是给一个协程访问数据的权限
 1.访问缓存
 2. 管理状态
**/

func main() {
	fmt.Println("go-channel run !")
	chan_demo()
}

func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

// 仅发送
func sendOnly(ch chan<- int) {
	for i := 1; i < 5; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	defer close(ch)
}

func chan_demo() {
	// 创建一个channel,3标识缓冲长度
	ch := make(chan int, 3)

	// 启动发送goroutine
	go sendOnly(ch)

	// 启动接受routine
	go receiveOnly(ch)

	// 使用select 进行多路复用

	// time.After()
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			// 500ms
			time.Sleep(500 * time.Millisecond)
		}
	}

}
