package main

import (
	"context"
	"fmt"
	"time"
)

/**
select与channel使用，select可以实现多个channel的收发数据
可以使用一个goroutine处理多个channel通信
语法类似switch 由case分支，default分支，只不过select的每个case后面跟的时channel的收发操作

select {
case channel_name <- value:
case value := <- channel
}

**/

func main() {

	// 创建多个channel
	chan1 := make(chan int, 10)
	chan2 := make(chan int, 10)
	chan3 := make(chan int, 10)

	for i := 0; i < 10; i++ {
		chan1 <- i
		chan2 <- i
		chan3 <- i
	}
	fmt.Println("channel has initlizal")

	// 一般启动一个上下文的context,用于为某些方法调用设置超时时间,当前线程阻塞主线程，防止未完全执行
	ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)
	var chan11 []int
	go func() {
		for {
			select {
			case v, ok := <-chan1:
				fmt.Println("case chan1 :", v, ok)
				chan11 = append(chan11, v)

				// 两秒后会输出一个空channel
			case <-ctx.Done():
				fmt.Println("context done")
				err := ctx.Err().Error()
				fmt.Println("error : ", err)
				return
			}
		}

	}()
	fmt.Println("chan11 : ", chan11)

}
