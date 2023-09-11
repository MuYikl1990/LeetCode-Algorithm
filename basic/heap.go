package basic

// 最小TopK
func minK(nums []int, k int) []int {
	if len(nums) <= k || k < 1 {
		return nums
	}
	res := make([]int, 0, k+1)
	for _, value := range nums {
		res = pushHeap(res, value)

		if len(res) > k {
			res = popHeap(res)
		}
	}
	return res
}

func pushHeap(heap []int, value int) []int {
	heap = append(heap, value)
	i := len(heap) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if (heap)[parent] < (heap)[i] {
			(heap)[parent], (heap)[i] = (heap)[i], (heap)[parent]
			i = parent
		} else {
			break
		}
	}
	return heap
}

func popHeap(heap []int) []int {
	head, i := 0, len(heap) - 1
	(heap)[head], (heap)[i] = (heap)[i], (heap)[head]
	heap = (heap)[:i]
	child := 2 * head + 1
	for child < len(heap) {
		if child + 1 < len(heap) && (heap)[child + 1] > (heap)[child] {
			child++
		}
		if (heap)[child] > (heap)[head] {
			(heap)[child], (heap)[head] = (heap)[head], (heap)[child]
			head = child
		} else {
			break
		}
	}
	return heap
}
