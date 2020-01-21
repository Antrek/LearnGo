package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct{
	taskid int
	createtime int64
}

type TaskQueue struct{
	TaskQueueChannel chan Task
}

func (this *TaskQueue) Init(){
	this.TaskQueueChannel = make(chan Task,10)
}

func (this *TaskQueue) Produce(){
	var task Task

	for {
		task.taskid = rand.Intn(100)
		task.createtime = time.Now().Unix()
		this.TaskQueueChannel<-task
		fmt.Println(task)
	}
}

func (this *TaskQueue)Consumer(){
	for {
		_ = <-this.TaskQueueChannel
		fmt.Println("已经消费一个任务")
		time.Sleep(time.Second)
	}

}

func(this *TaskQueue) xx() {
	fmt.Println("xx")
}

func main() {
	var taskqueue TaskQueue
	taskqueue.Init()
	go taskqueue.Produce()
	go taskqueue.Consumer()

	time.Sleep(time.Hour)
}


