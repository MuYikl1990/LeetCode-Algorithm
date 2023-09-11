package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	courses := make([]int, numCourses)
	mapCourse := make(map[int][]int)

	for _, arr := range prerequisites {
		courses[arr[0]] += 1
		mapCourse[arr[1]] = append(mapCourse[arr[1]], arr[0])
	}

	var queue []int
	for key, value := range courses {
		if value == 0 {
			queue = append(queue, key)
		}
	}

	count := 0
	for len(queue) > 0 {
		cur := queue[0]
		count += 1
		list, ok := mapCourse[cur]; if ok {
			for _, val := range list {
				courses[val] -= 1
				if courses[val] == 0 {
					queue = append(queue, val)
				}
			}
		}
		queue = queue[1:]
	}
	return numCourses == count
}

func main() {
	numCourses := 3
	prerequisites := [][]int{{0,1}, {1,2}, {0,2}, {2,1}}
	res := canFinish(numCourses, prerequisites)
	fmt.Println(res)
}
