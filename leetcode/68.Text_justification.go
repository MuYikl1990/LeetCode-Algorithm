package main

import "fmt"

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	start, count := 0, 0
	for i := 0; i < len(words); i++ {
		count += len(words[i]) + 1
		if i == len(words) - 1 || count + len(words[i + 1]) > maxWidth {
			res = append(res, fullText(words, start, i, maxWidth, i == len(words) - 1))
			count = 0
			start = i + 1
		}
	}
	return res
}

func fullText(words []string, start, end, maxWidth int, isLast bool) string {
	res := ""
	wordCount := end - start + 1
	space := maxWidth + 1 - wordCount
	for i := start; i <= end; i++ {
		space -= len(words[i])
	}

	avg, extra := 1, 0
	if wordCount > 1 {
		avg, extra = space / (wordCount - 1), space % (wordCount - 1)
	}

	for i := start; i < end; i++ {
		res += words[i] + " "
		if !isLast {
			for k := 0; k < avg; k++ {
				res += " "
			}
			if extra > 0 {
				res += " "
				extra--
			}
		}
	}
	res += words[end]
	for len(res) != maxWidth {
		res += " "
	}
	fmt.Println(len(res))
	return res
}

func main() {
	words, maxWidth := []string{"Science","is","what","we","understand","well","enough","to","explain","to","a",
		"computer.","Art","is","everything","else","we","do"}, 20
	res := fullJustify(words, maxWidth)
	fmt.Println(res)
}
