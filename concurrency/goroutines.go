package concurrency

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second) // 模拟耗时任务
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func RunGoroutinesDemo() {
	// 创建jobs和results通道
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动3个worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
