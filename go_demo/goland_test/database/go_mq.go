package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数据生产者
func producer2(header string, channel chan<- string) {
	// 无限循环, 不停地生产数据
	for {
		// 将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		// 等待1秒
		time.Sleep(time.Second)
	}
}

// 数据消费者
func customer(channel <-chan string) {
	// 不停地获取数据
	for {
		// 从通道中取出数据, 此处会阻塞直到信道中返回数据
		message := <-channel
		// 打印数据
		fmt.Println(message)
	}
}

func main() {
	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 创建producer()函数的并发goroutine
	go producer2("cat", channel)
	go producer2("dog", channel)
	// 数据消费函数
	customer(channel)
}