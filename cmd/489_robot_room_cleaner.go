package main

import "fmt"

//lint:ignore U1000 unused
func cleanRoom(robot *Robot) {
	cleanRoomDFS(robot, 0, 0, 0, map[[2]int]bool{})
}

func cleanRoomDFS(r *Robot, i, j int, direction int, visited map[[2]int]bool) {
	if visited[[2]int{i, j}] {
		return
	}
	r.Clean()
	visited[[2]int{i, j}] = true
	for n := 0; n < 4; n++ {
		if r.Move() {
			row := i
			col := j
			switch direction {
			case 0:
				row = i - 1
			case 90:
				col = j + 1
			case 180:
				row = i + 1
			case 270:
				col = j - 1
			}

			cleanRoomDFS(r, row, col, direction, visited)
			r.TurnLeft()
			r.TurnLeft()
			r.Move()
			r.TurnRight()
			r.TurnRight()
		}
		r.TurnRight()
		direction += 90
		direction = direction % 360
	}
}

type Robot struct {
	direction int
	room      [][]int
	result    [][]int
	x         int
	y         int
}

func (robot *Robot) Move() bool {
	x := 0
	y := 0
	switch robot.direction {
	case 0:
		y = -1
	case 90:
		x = 1
	case 180:
		y = 1
	case 270:
		x = -1
	}
	newX := robot.x + x
	newY := robot.y + y

	outOfbounds := newX < 0 || newY < 0 || newX > len(robot.room)-1 || newY > len(robot.room[0])-1
	if outOfbounds || robot.room[newX][newY] == 0 {
		return false
	}
	robot.x = newX
	robot.y = newY
	fmt.Println("Robot at: ", robot.x, " , ", robot.y)

	return true
}

func (robot *Robot) TurnLeft() {
	if robot.direction == 0 {
		robot.direction = 270
	} else {
		robot.direction -= 90
	}
}
func (robot *Robot) TurnRight() {
	robot.direction += 90
	if robot.direction == 360 {
		robot.direction = 0
	}
}
func (robot *Robot) Clean() {
	robot.result[robot.x][robot.y] = 1
}
