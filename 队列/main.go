package main

import (
	"errors"
	"fmt"
)

func main() {
	this := InitPriorityQueue(2)
	fmt.Printf("111111111111=%v\n", this.Arr)
	this.EnQueue(3)
	fmt.Printf("2222222222222=%v\n", this.Arr)
	this.EnQueue(8)
	fmt.Printf("333333333333333=%v\n", this.Arr)
	this.EnQueue(1)
	fmt.Printf("4444444444444444=%v\n", this.Arr)
	this.EnQueue(9)
	fmt.Printf("5555555555555555=%v\n", this.Arr)
	this.DeQueue()
	fmt.Printf("666666=%v\n", this.Arr)
	this.DeQueue()
	fmt.Printf("7777777=%v\n", this.Arr)
	this.DeQueue()
	fmt.Printf("888888=%v\n", this.Arr)
	this.EnQueue(10)
	fmt.Printf("333333333333333=%v\n", this.Arr)
	this.EnQueue(11)
	fmt.Printf("333333333333333=%v\n", this.Arr)
	this.EnQueue(12)
	fmt.Printf("333333333333333=%v\n", this.Arr)
	this.EnQueue(13)
	fmt.Printf("333333333333333=%v\n", this.Arr)
	this.EnQueue(14)
	fmt.Printf("333333333333333=%v\n", this.Arr)
}

// start-------------------------双端循环队列------------------------start
type MyCircularDeque struct {
	Arr      []int
	Head     int
	Rear     int
	Capacity int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	arr := make([]int, k+1)
	return MyCircularDeque{
		Arr:      arr,
		Head:     0,
		Rear:     0,
		Capacity: k + 1,
	}
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.Head == this.Rear

}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.Rear+1 == this.Head

}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	front := (this.Head - 1 + this.Capacity) % this.Capacity
	this.Arr[front] = value
	this.Head = front
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Arr[this.Rear] = value
	this.Rear = (this.Rear + 1 + this.Capacity) % this.Capacity
	return true

}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = (this.Head + 1 + this.Capacity) % this.Capacity
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.Rear = (this.Rear - 1 + this.Capacity) % this.Capacity
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Arr[this.Head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Arr[(this.Rear-1+this.Capacity)%this.Capacity]
}

// end----------------------双端循环队列---------------------------end

// start---------------------优先队列---------------------------start

type PriorityQueue struct {
	Arr  []int
	Size int
}

func InitPriorityQueue(size int) *PriorityQueue {
	arr := make([]int, size)
	return &PriorityQueue{
		Arr:  arr,
		Size: 0,
	}
}
func (this *PriorityQueue) EnQueue(value int) {
	if len(this.Arr) == this.Size {
		this.Resize()
	}
	this.Arr[this.Size] = value
	this.Size = this.Size + 1
	this.upAdjust()
}
func (this *PriorityQueue) DeQueue() (int, error) {
	if this.Size == 0 {
		return 0, errors.New("对列为空")
	}
	value := this.Arr[0]
	i := this.Size - 1
	this.Arr[0] = this.Arr[i]
	this.Arr[i] = 0
	this.Size--
	this.downAdjust()
	return value, nil

}
func (this *PriorityQueue) Resize() {
	expandArr := make([]int, this.Size)
	this.Arr = append(this.Arr, expandArr...)
}

// 下沉节点:删除的时候下沉
func (this *PriorityQueue) downAdjust() {
	parentIndex := 0
	temp := this.Arr[0]
	childIndex := 2*parentIndex + 1
	lenth := this.Size
	for childIndex < lenth {
		if childIndex+1 < lenth && this.Arr[childIndex+1] > this.Arr[childIndex] {
			childIndex = childIndex + 1
		}
		if temp > this.Arr[childIndex] {
			break
		}
		this.Arr[parentIndex] = this.Arr[childIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}
	this.Arr[parentIndex] = temp
}

// 上浮节点,添加节点的时候上浮
func (this *PriorityQueue) upAdjust() {
	childIndex := this.Size - 1

	temp := this.Arr[childIndex]
	parentIndex := (childIndex - 1) / 2
	for childIndex > 0 && temp > this.Arr[parentIndex] {
		this.Arr[childIndex] = this.Arr[parentIndex]
		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
	this.Arr[childIndex] = temp
}

// end----------------------优先队列---------------------------end
