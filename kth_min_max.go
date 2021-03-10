package main

import "fmt"

func main() {

	var a = []int{11, 14, 7, 8, -10}
	 min := findKthLargest(a, 3)
//	fmt.Println("Min: ", min)
	fmt.Println("Max: ", min)

}

func findKthLargest(nums []int, k int) int {
    if len(nums) < k{
        return 0
    }

    heap := make([]int, 0)
    for _, num := range nums{
        if len(heap) < k{
            heap = append(heap, num)
            if len(heap) == k{
                buildHeap(heap)
            }
            continue
        }

        if num > heap[0]{
            heap[0] = num
            heapHelp(heap, k, 0)
        }
    }

    return heap[0]
}

func buildHeap(heap []int){
    size := len(heap)

    for i := size - 1; i >= 0; i--{
        heapHelp(heap, size, i)
    }

    return
}

func heapHelp(heap []int, size, curRoot int){
    left := 2*curRoot+1
    right := 2*curRoot+2
    smallest := curRoot

    if left < size && heap[left] < heap[smallest]{
        smallest = left
    }

    if right < size && heap[right] < heap[smallest]{
        smallest = right
    }

    if smallest != curRoot{
        heap[smallest], heap[curRoot] = heap[curRoot], heap[smallest]
        heapHelp(heap, size, smallest)
    }

    return
}
