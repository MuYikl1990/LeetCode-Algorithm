package main

import "fmt"

func removeComments(source []string) []string {
	inComment := false
	var res []string
	var newLine []byte

	for _, words := range source {
		if !inComment {
			newLine = []byte{}
		}
		for i := 0; i < len(words); {
			if len(words) - i > 1 && words[i:i + 2] == "/*" && !inComment {
				inComment = true
				i += 1
			} else if len(words) - i > 1 && words[i:i + 2] == "*/" && inComment {
				inComment = false
				i += 1
			} else if len(words) - i > 1 && words[i:i + 2] == "//" && !inComment {
				i += 1
				break
			} else if !inComment {
				newLine = append(newLine, words[i])
			}
			i++
		}
		if len(newLine) != 0 && !inComment {
			res = append(res, string(newLine))
		}
	}
	return res
}

func main() {
	source := []string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}
	res := removeComments(source)
	fmt.Println(res)
}
