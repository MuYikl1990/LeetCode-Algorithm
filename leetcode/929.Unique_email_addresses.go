package main

import "fmt"

func numUniqueEmails(emails []string) int {
	res := 0
	eMap := make(map[string]struct{})
	for _, email := range emails {
		var tmp []byte
		i := 0
		for i < len(email) {
			if email[i] == '@' {
				break
			}
			if email[i] == '.' {
				i++
				continue
			}
			if email[i] == '+' {
				for email[i] != '@' {
					i++
				}
				break
			}
			tmp = append(tmp, email[i])
			i++
		}
		tmp = append(tmp, []byte(email[i:])...)
		if _, ok := eMap[string(tmp)]; !ok {
			res++
			eMap[string(tmp)] = struct{}{}
		}
	}
	return res
}

func main() {
	emails := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	res := numUniqueEmails(emails)
	fmt.Println(res)
}
