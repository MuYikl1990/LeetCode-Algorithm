package basic

// 冒泡排序
func BubbleSort(nums []int) []int {
	length := len(nums)

	if length < 2 {
		return nums
	}

	for i := 0; i < length - 1; i++ {
		flag := false

		for j := 0; j < length - i - 1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nums
}

// 选择排序
func SelectSort(nums []int) []int {
	length := len(nums)

	if length < 2 {
		return nums
	}

	for i, _ := range nums {
		min := nums[i]
		index := i
		for j := i; j < length; j++ {
			if nums[j] < min {
				min = nums[j]
				index = j
			}
		}
		nums[i], nums[index] = nums[index], nums[i]
	}
	return nums
}

// 插入排序
func InsertSort(nums []int) []int {
	length := len(nums)

	if length < 2 {
		return nums
	}

	for i, _ := range nums {
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}
	return nums
}

// 归并排序
func MergeSort(nums []int) []int {
	length := len(nums)

	if length < 2 {
		return nums
	}

	mid := length / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	lIndex, rIndex := 0, 0
	var res []int
	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] <= right[rIndex] {
			res = append(res, left[lIndex])
			lIndex++
		} else {
			res = append(res, right[rIndex])
			rIndex++
		}
	}

	res = append(res, left[lIndex:]...)
	res = append(res, right[rIndex:]...)
	return res
}

// 快速排序
func QuickSort(nums []int) []int {
	partition(nums, 0, len(nums)-1)
	return nums
}

func partition(nums []int, start, end int) {
	if start >= end {
		return
	}

	pivot := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] <= pivot {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]

	partition(nums, start, i-1)
	partition(nums, i+1, end)
}

// 堆排序
func HeapSort(nums []int) []int {
	length := len(nums)
	buildMaxHeap(nums, length)
	for i := length - 1; i >= 0; i-- {
		swapNode(nums, 0, i)
		length--
		heapify(nums, 0, length)
	}
	return nums
}

func buildMaxHeap(nums []int, length int) {
	for i := length / 2; i >= 0; i-- {
		heapify(nums, i, length)
	}
}

func heapify(nums []int, start, length int) {
	left := start * 2 + 1
	right := start * 2 + 2
	large := start
	if left < length && nums[left] > nums[large] {
		large = left
	}
	if right < length && nums[right] > nums[large] {
		large = right
	}
	if large != start {
		swapNode(nums, start, large)
		heapify(nums, large, length)
	}
}

func swapNode(nums []int, start, end int) {
	nums[start], nums[end] = nums[end], nums[start]
}

// 堆排序迭代
func HeapSortWithIteration(nums []int) []int {
	n := len(nums)
	buildHeap(nums, n)
	k := n
	for k > 0 {
		swap1(nums, 0, k - 1)
		k--
		heapify1(nums, k, 0)
	}
	return nums
}

func buildHeap(nums []int, n int) {
	for i := (n - 1) / 2; i >= 0; i-- {
		heapify1(nums, n, i)
	}
}

func heapify1(nums []int, n, i int) {
	for {
		idx := i
		if i * 2 + 1 < n && nums[i * 2 + 1] > nums[i] {
			idx = i * 2 + 1
		}
		if i * 2 + 2 < n && nums[i * 2 + 2] > nums[idx] {
			idx = i * 2 + 2
		}
		if idx == i {
			break
		}
		swap1(nums, i, idx)
		i = idx
	}
}

func swap1(nums []int, start, end int) {
	nums[start], nums[end] = nums[end], nums[start]
}
