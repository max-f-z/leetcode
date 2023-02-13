package main

func countOdds(low int, high int) int {
	lowOdds, highOdds := 0, 0
	lowOdds = low / 2

	if high%2 == 0 {
		highOdds = high / 2
	} else {
		highOdds = high/2 + 1
	}

	return highOdds - lowOdds
}
