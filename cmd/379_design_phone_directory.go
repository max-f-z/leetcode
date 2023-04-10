package main

type PhoneDirectory struct {
	number []int
}

func Constructor379(maxNumbers int) PhoneDirectory {
	var number []int
	for i := 0; i < maxNumbers; i++ {
		number = append(number, i)
	}
	return PhoneDirectory{number: number}
}

//lint:ignore ST1006 this
func (this *PhoneDirectory) Get() int {
	if len(this.number) == 0 {
		return -1
	}

	tmp := this.number[0]
	this.number = this.number[1:]
	return tmp
}

//lint:ignore ST1006 this
func (this *PhoneDirectory) Check(number int) bool {
	for _, v := range this.number {
		if v == number {
			return true
		}
	}
	return false
}

//lint:ignore ST1006 this
func (this *PhoneDirectory) Release(number int) {
	flag := false
	for _, v := range this.number {
		if v == number {
			flag = true
		}
	}
	if !flag {
		this.number = append(this.number, number)

	}
}
