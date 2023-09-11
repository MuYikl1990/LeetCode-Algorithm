package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if queryIP == "" {
		return "Neither"
	}

	list4 := strings.Split(queryIP, ".")
	if len(list4) == 4 {
		for _, ip4 := range list4 {
			if len(ip4) > 1 && ip4[0] == '0' {
				return "Neither"
			}
			number, err := strconv.Atoi(ip4)
			if err != nil {
				return "Neither"
			}
			if number < 0 || number > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	} else {
		list6 := strings.Split(queryIP, ":")
		if len(list6) == 8 {
			for _, ip6 := range list6 {
				if len(ip6) < 1 || len(ip6) > 4 {
					return "Neither"
				}
				if _, err := strconv.ParseUint(ip6, 16, 64); err != nil {
					return "Neither"
				}
			}
			return "IPv6"
		} else {
			return "Neither"
		}
	}
}

func main() {
	queryIP := "2001:0db8:85a3:8f:0:8A2E:037E:7334"
	res := validIPAddress(queryIP)
	fmt.Println(res)
}
