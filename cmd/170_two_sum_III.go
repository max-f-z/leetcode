package main

type TwoSum struct {
	times map[int]int
	data  []int
}

/** Initialize your data structure here. */
func Constructor1() TwoSum {
	ts := TwoSum{}
	ts.times = map[int]int{}
	ts.data = []int{}
	return ts
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.times[number]++
	this.data = append(this.data, number)
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for _, v := range this.data {
		remain := value - v
		if (remain == v && this.times[remain] > 1) || (remain != v && this.times[remain] > 0) {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
