package main

import (
	"fmt"
	"runtime"
)

func main(){
	c := make(chan int)
	quit := make(chan int)
	go func(){
		for i := 0; i < 10; i++ {
			fmt.Printf("%d ", <-c )
		}
		quit <- 1
	}()
	testMulti(c, quit)
}

func testMulti(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main2() {
	//var f chan<- int
	f := make(chan int)

	f <- 1

	v := <-f

	fmt.Println(v)

	//buf := make(chan int)
	//flg := make(chan int)
	//go producer(buf)
	//go consumer(buf, flg)
	//v := <-flg // 等待接收完成
}


//func producer(c chan int) {
func producer(c chan<- int) {
	defer close(c) // 关闭 channel
	for i := 0; i < 10; i++ {
		c <- i // 阻塞，直到数据被消费者取走后，才能发送下一条数据
	}
}

//func consumer(c, f chan int) {
func consumer(c <-chan int, f chan<- int) {
	for {
		if v, ok := <-c; ok {
			fmt.Println(v) // 阻塞， 直到生产者放入数据后继续读取数据
		} else {
			break
		}
	}
	f <- 1 // 发送数据，通知 main函数已接收完成。
}

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Print("hello ")
		runtime.Gosched()
	}
}

func sayWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("world")
		runtime.Gosched()
	}
}
