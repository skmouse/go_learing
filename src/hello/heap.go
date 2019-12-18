package main

import (
	"fmt"
	"math"
)

type MaxHeap struct {
	data []  int
}

type MinHeap struct {
	data []  int
}

/**
 * init heap
 */
func NewMaxHeap() *MaxHeap {
	h := &MaxHeap{data: []int{math.MaxInt64}}
	return h
}

/**
 * insert max heap
 */
func (H *MaxHeap) insert(val int)  {
	H.data = append(H.data, val)
	i := len(H.data) -1

	for ;H.data[i/2] < val ; i/=2 {
		H.data[i] = H.data[i/2]
	}

	H.data[i] = val
}

/**
 * delete max heap
 */
func (H *MaxHeap) DelateMaxHeap() int {
	var parent, child int
	maxItem := H.data[1]

	temp := H.data[len(H.data) -1]

	for parent=1; parent*2 <= len(H.data) -1; parent = child  {
		child = parent*2
		if (child != len(H.data) -1) && (H.data[child] < H.data[child+1]) {
			child ++
		}

		if temp >= H.data[child] {
			break
		} else {
			H.data[parent] = H.data[child]
		}
	}

	H.data[parent] = temp
	return maxItem
}

/**
 * init MinHeap
 */
func NewMinHeap() *MinHeap  {
	h := &MinHeap{data: []int{math.MinInt64}}
	return h
}

/**
 * insert MinHeap
 */
func (H* MinHeap) insertMinHeap(val int) {
	H.data = append(H.data, val)
	i := len(H.data) -1

	for ; H.data[i/2] > val ; i/=2 {
		H.data[i] = H.data[i/2]
	}
	H.data[i] = val
}

/**
 * delete MinHeap
 */
func (H *MinHeap) deleteMinHeap() int {
	var parent, child int
	MinItem :=H.data[1]
	temp := H.data[len(H.data) -1]
	length := len(H.data) -1
	for parent=1; parent*2 < length ; parent = child  {
		child = parent*2
		if H.data[child] != length && H.data[child] > H.data[child+1] {
			child++
		}

		if temp <= H.data[child] {
			break
		} else {
			H.data[parent] = H.data[child]
		}
	}

	H.data[parent] = temp
	return MinItem
}

func (H *MaxHeap) Max() int {
	if len(H.data) >1 {
		return H.data[1]
	}
	return  -1
}

func (H *MinHeap) Min() int  {
	if len(H.data) >1 {
		return H.data[1]
	}
	return  -1
}

func main()  {
	heap := NewMaxHeap()
	heap.insert(44)
	heap.insert(25)
	heap.insert(31)
	heap.insert(18)
	heap.insert(10)
	fmt.Println(heap.data)

	fmt.Print(heap.DelateMaxHeap(), " ")
	fmt.Println(heap.data)



	fmt.Println("insert minheap", " ")
	heapMin := NewMinHeap()
	heapMin.insertMinHeap(44)
	heapMin.insertMinHeap(25)
	heapMin.insertMinHeap(31)
	heapMin.insertMinHeap(18)
	heapMin.insertMinHeap(10)
	fmt.Println(heapMin.data)


	fmt.Print("delete minheap", " ")

	fmt.Print(heapMin.deleteMinHeap())
	fmt.Print(heapMin.data)
}