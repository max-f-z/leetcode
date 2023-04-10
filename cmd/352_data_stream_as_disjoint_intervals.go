package main

type SummaryRanges struct {
	ranges [][]int
}

func Constructor352() SummaryRanges {
	return SummaryRanges{
		ranges: [][]int{},
	}
}

//lint:ignore ST1006 this
func (this *SummaryRanges) AddNum(value int) {
	if len(this.ranges) == 0 {
		this.ranges = [][]int{{value, value}}
		return
	}

	low, high := 0, len(this.ranges)-1
	for low <= high {
		mid := (low + high) / 2

		if this.ranges[mid][0] <= value && this.ranges[mid][1] >= value {
			return
		}

		if this.ranges[mid][1]+1 == value {
			this.ranges[mid][1] = value

			if mid < len(this.ranges)-1 && this.ranges[mid+1][0] == value+1 {
				this.ranges[mid+1][0] = this.ranges[mid][0]
				this.ranges = append(this.ranges[:mid], this.ranges[mid+1:]...)
			}
			return
		}

		if this.ranges[mid][0]-1 == value {
			this.ranges[mid][0] = value

			if mid >= 1 && this.ranges[mid-1][1] == value-1 {
				this.ranges[mid-1][1] = this.ranges[mid][1]
				this.ranges = append(this.ranges[:mid], this.ranges[mid+1:]...)
			}
			return
		}

		if this.ranges[mid][1] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	this.ranges = append(this.ranges[:low], append([][]int{{value, value}}, this.ranges[low:]...)...)
}

//lint:ignore ST1006 this
func (this *SummaryRanges) GetIntervals() [][]int {
	return this.ranges
}
