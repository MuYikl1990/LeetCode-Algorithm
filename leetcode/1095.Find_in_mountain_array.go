package main

import "fmt"

type MountainArray struct {
	arr []int
}

func (this *MountainArray) get(index int) int {
	return this.arr[index]
}
func (this *MountainArray) length() int {
	return len(this.arr)
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	length := mountainArr.length()
	peak := findInMountainPeak(mountainArr, 0, length - 1)
	res := findInMountainLeft(mountainArr, 0, peak, target)
	if res != -1 {
		return res
	}
	return findInMountainRight(mountainArr, peak + 1, length - 1, target)
}

func findInMountainPeak(mountainArr *MountainArray, left, right int) int {
	for left < right {
		mid := left + (right - left) / 2
		if mountainArr.get(mid) < mountainArr.get(mid + 1) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func findInMountainLeft(mountainArr *MountainArray, left, right, target int) int {
	for left < right {
		mid := left + (right - left) / 2
		if mountainArr.get(mid) < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if mountainArr.get(left) == target {
		return left
	}
	return -1
}

func findInMountainRight(mountainArr *MountainArray, left, right, target int) int {
	for left < right {
		mid := left + (right - left + 1) / 2
		if mountainArr.get(mid) < target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if mountainArr.get(left) == target {
		return left
	}
	return -1
}

func main() {
	mountainArr, target := &MountainArray{arr: []int{1,2,3,4,5,3,1}}, 3
	res := findInMountainArray(target, mountainArr)
	fmt.Println(res)
}
