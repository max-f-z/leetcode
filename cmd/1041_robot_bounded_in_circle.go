package main

//lint:ignore U1000 unused
func isRobotBounded(instructions string) bool {

	x, y := 0, 0

	direction := 0

	for i := 0; i < len(instructions); i++ {
		if instructions[i] == 'G' {
			switch direction {
			case 0:
				y += 1
			case 1:
				x -= 1
			case 2:
				y -= 1
			case 3:
				x += 1
			}
		} else if instructions[i] == 'L' {
			direction = (direction + 1) % 4
		} else if instructions[i] == 'R' {
			direction = (direction - 1 + 4) % 4
		}
	}

	if direction != 0 {
		return true
	}

	if x == 0 && y == 0 {
		return true
	}

	return false
}
