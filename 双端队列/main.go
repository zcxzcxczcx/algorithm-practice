package main

func main() {

}

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
	front := (this.Head + 1 + this.Capacity) % this.Capacity
	this.Head = front
	return true

}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	last := (this.Rear - 1 + this.Capacity) % this.Capacity
	this.Rear = last
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

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.Head == this.Rear

}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {

	return (this.Rear+1)%this.Capacity == this.Head
}
