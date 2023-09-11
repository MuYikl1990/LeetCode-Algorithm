package main

import (
	"fmt"
	"strconv"
)

func getFolderNames(names []string) []string {
	var res []string
	folders := make(map[string]int)

	for _, name := range names {
		if count, ok := folders[name]; ok {
			for i := count; ; i++ {
				str := name + "(" + strconv.Itoa(i) + ")"
				if _, exist := folders[str]; !exist {
					folders[name] = i + 1
					res = append(res, str)
					folders[str] = 1
					break
				}
			}
		} else {
			folders[name] = 1
			res = append(res, name)
		}
	}
	return res
}

func main() {
	names := []string{"onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece","onepiece(1)"}
	res := getFolderNames(names)
	fmt.Println(res)
}
