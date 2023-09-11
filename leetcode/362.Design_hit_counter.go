package main

import "fmt"

type HitCounter struct {
	queue []int
}

func Constructor362() HitCounter {
	return HitCounter{}
}

func (h *HitCounter) hit(timestamp int) {
	h.queue = append(h.queue, timestamp)
}

func (h *HitCounter) getHits(timestamp int) int {
	for len(h.queue) != 0 && timestamp - (h.queue)[0]>= 300 {
		h.queue = (h.queue)[1:]
	}
	return len(h.queue)
}

func main() {
	h := Constructor362()
	h.hit(1)
	h.hit(3)
	res1 := h.getHits(7)
	res2 := h.getHits(303)
	fmt.Println(res1)
	fmt.Println(res2)
}
