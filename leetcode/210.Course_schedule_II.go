package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	var res []int
	if len(prerequisites) == 0 {
		for i := 0; i < numCourses; i++ {
			res = append(res, i)
		}
		return res
	}

	list := make([]int, numCourses)
	courseMap := make(map[int][]int)

	for _, prerequisite := range prerequisites {
		list[prerequisite[0]] += 1
		courseMap[prerequisite[1]] = append(courseMap[prerequisite[1]], prerequisite[0])
	}

	var queue []int
	for i, val := range list {
		if val == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		count++
		cur := queue[0]
		res = append(res, cur)
		tmp := courseMap[cur]
		for _, val := range tmp {
			list[val]--
			if list[val] == 0 {
				queue = append(queue, val)
			}
		}
		queue = queue[1:]
	}
	if count != numCourses {
		return []int{}
	}
	return res
}

func main() {
	numCourses, prerequisites := 4, [][]int{{1,0}, {2,0}, {3,1}, {3,2}, {0,1}}
	res := findOrder(numCourses, prerequisites)
	fmt.Println(res)
}
