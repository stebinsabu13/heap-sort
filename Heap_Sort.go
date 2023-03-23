package main

import "fmt"

func heapSort(arr []int) {
	n := len(arr)
	
	//loop to make the array with max heap properties
	for i := (len(arr) / 2) - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	
	//loop to make the array sorted
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

//heapify function to maintain the heap properties

func heapify(arr []int, n, i int) {
	largest := i
	lc := lchild(i)
	rc := rchild(i)
	if lc < n && arr[largest] < arr[lc] {
		largest = lc
	}
	if rc < n && arr[largest] < arr[rc] {
		largest = rc
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
func lchild(i int) int {
	return (i * 2) + 1
}
func rchild(i int) int {
	return (i * 2) + 2
}
func main() {
	arr := []int{10, 20, 15, 12, 40, 25, 18}
	heapSort(arr)
	fmt.Println(arr)
}
