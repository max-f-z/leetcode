package main

// https://haogroot.com/2020/09/01/monotonic-stack-leetcode/

type StockSpanner struct {
	elements []*StockSpannerElem
	count    int
}

type StockSpannerElem struct {
	index int
	value int
}

func Constructor901() StockSpanner {
	return StockSpanner{
		elements: make([]*StockSpannerElem, 0),
		count:    0,
	}
}

//lint:ignore ST1006 this
func (this *StockSpanner) Next(price int) int {
	this.count++

	if len(this.elements) == 0 {
		this.elements = append(this.elements, &StockSpannerElem{
			index: 1,
			value: price,
		})
		return 1
	}

	for i := len(this.elements) - 1; i >= 0; i-- {
		if price >= this.elements[i].value {
			continue
		} else {
			if i < len(this.elements)-1 {
				this.elements = this.elements[:i+1]
			}
			this.elements = append(this.elements, &StockSpannerElem{
				index: this.count,
				value: price,
			})
			return this.count - this.elements[i].index
		}
	}

	this.elements = []*StockSpannerElem{
		{
			index: this.count,
			value: price,
		},
	}
	return this.count
}
