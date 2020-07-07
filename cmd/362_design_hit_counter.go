package main

type HitCounter struct {
	data []int
}

func Constructor362() HitCounter {
	hc := HitCounter{}
	hc.data = []int{}
	return hc
}

func (this *HitCounter) Hit(timestamp int) {
	idx := -1
	for i := 0; i < len(this.data); i++ {
		if this.data[i] < timestamp-300 {
			idx = i
		} else {
			break
		}
	}
	if idx != -1 {
		this.data = this.data[idx+1:]
	}

	this.data = append(this.data, timestamp)
}

func (this *HitCounter) GetHits(timestamp int) int {
	idx := -1
	for i := 0; i < len(this.data); i++ {
		if this.data[i] <= timestamp-300 {
			idx = i
		} else {
			break
		}
	}
	if idx != -1 {
		this.data = this.data[idx+1:]
	}

	return len(this.data)
}
