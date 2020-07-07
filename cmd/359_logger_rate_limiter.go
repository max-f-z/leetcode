package main

type Logger struct {
	data map[string]int
}

func Constructor359() Logger {
	l := Logger{}
	l.data = map[string]int{}
	return l
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if _, ok := this.data[message]; !ok {
		this.data[message] = timestamp
		return true
	} else {
		tmp := this.data[message]
		if timestamp-tmp >= 10 {
			this.data[message] = timestamp
			return true
		} else {
			return false
		}
	}
}
