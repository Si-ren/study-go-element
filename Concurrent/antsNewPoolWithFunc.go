package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

const (
	DataSize    = 10000
	DataPerTask = 100
)

type Task struct {
	index int
	nums  []int
	sum   int
	wg    *sync.WaitGroup
}

func (t *Task) Do() {
	for _, num := range t.nums {
		t.sum += num
	}

	t.wg.Done()
}

func taskFunc(data interface{}) {
	task := data.(*Task)
	task.Do()
	fmt.Printf("task:%d sum:%d\n", task.index, task.sum)
}

func main() {
	// 这里建个日志
	z, err := zap.NewProduction()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// 创建协程池,ants.NewPoolWithFunc可以创建传参的func协程池
	// 如果是只是协程池,用ants.NewPool(),然后用传入func(),用闭包的方式传入参数
	p, err := ants.NewPoolWithFunc(5000, taskFunc)
	if err != nil {
		z.Error(err.Error())
	}
	// main结束释放协程池
	defer p.Release()

	// 创建slice,并且随机塞入数值
	nums := make([]int, DataSize, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}

	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	tasks := make([]*Task, 0, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		task := &Task{
			index: i + 1,
			nums:  nums[i*DataPerTask : (i+1)*DataPerTask],
			wg:    &wg,
		}
		// 这边slice不够了会自动扩容
		tasks = append(tasks, task)
		// 把task切片放进协程池,相当于给协程池中的协程传参数
		p.Invoke(task)
	}

	// 等待所有都处理完后
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())

	var sum int
	for _, task := range tasks {
		sum += task.sum
	}

	var expect int
	for _, num := range nums {
		expect += num
	}

	fmt.Printf("finish all tasks, result is %d expect:%d\n", sum, expect)
}
