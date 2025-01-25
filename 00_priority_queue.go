package main

import (
	"container/heap"
	"fmt"
)

// Define a type that implements heap.Interface
type MinHeap []int

// Implement heap.Interface
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes the smallest element from the heap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	h := &MinHeap{5, 4, 3, 2, 1, 0}

	heap.Push(h, 3)
	heap.Push(h, 8)
	//
	fmt.Println("Min-heap", h)
	heap.Init(h)

	fmt.Println("Min-heap value", *h)

	heap.Remove(h, 2)

	fmt.Println("Pop:", heap.Pop(h))
	fmt.Println("Min-Heap after pop:", *h)

	//for h.Len() > 0 {
	//	fmt.Println("Pop:", heap.Pop(h))
	//}

}

//Internally heap algorithm uses slice and has all operation on that slice - it uses below rules - use 0 based indexing
//1. root element index = (i-1)/2
//2. left child = i*2 +1
//3. right child = i*2 + 2

// min heap is complete binary tree with parent node less than children
// max heap is complete binary tree with parent node is greater than its children
