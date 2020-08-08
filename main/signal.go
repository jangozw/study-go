package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("信号测试...")

	/*for {
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
	}*/

	/**
	上述代码监听了如下信号：
	SIGINT: Ctrl-C 发送 INT signal (SIGINT)，通常导致进程结束
	SIGTERM: 程序结束(terminate)信号
	有两种信号不能被拦截和处理: SIGKILL 和 SIGSTOP
	*/

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan
		/* 可以在下面定义一些资源回收，代码清理工作 */
		fmt.Println("\n收到信号：进程结束")
	}()

	fmt.Println("9999")
	time.Sleep(10 * time.Second)
}
