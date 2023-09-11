package main

import "fmt"

type stock struct {
	price int
	day   int
}

type StockSpanner struct {
	span []stock
}


func Constructor901() StockSpanner {
	return StockSpanner{[]stock{}}
}


func (this *StockSpanner) Next(price int) int {
	res := 1
	for len(this.span) > 0 {
		cur := this.span[len(this.span) - 1]
		if cur.price <= price {
			res += cur.day
			this.span = this.span[:len(this.span) - 1]
		} else {
			break
		}
	}
	this.span = append(this.span, stock{price, res})
	return res
}

func main() {
	price := 100
	obj := Constructor901()
	res := obj.Next(price)
	fmt.Println(res)
}
