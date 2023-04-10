package main

import (
	"container/list"
)

type SnakeGame struct {
	body    *list.List
	food    [][]int
	board   [][]int
	end     *list.Element
	h       int
	w       int
	foodIdx int
}

/*
  - Initialize your data structure here.
    @param width - screen width
    @param height - screen height
    @param food - A list of food positions
    E.g food = [[1,1], [1,0]] means the first food is positioned at [1,1], the second is at [1,0].
*/
func Constructor353(width int, height int, food [][]int) SnakeGame {
	sg := SnakeGame{}
	sg.h = height
	sg.w = width
	sg.foodIdx = 0
	sg.food = food
	sg.body = list.New()
	sg.body.PushFront([]int{0, 0})
	sg.board = make([][]int, height)
	for i := 0; i < height; i++ {
		sg.board[i] = make([]int, width)
	}
	sg.board[0][0] = 1

	sg.end = sg.body.Front()

	return sg
}

/*
  - Moves the snake.
    @param direction - 'U' = Up, 'L' = Left, 'R' = Right, 'D' = Down
    @return The game's score after the move. Return -1 if game over.
    Game over when snake crosses the screen boundary or bites its body.
*/
//lint:ignore ST1006 this
func (this *SnakeGame) Move(direction string) int {
	head, _ := this.body.Front().Value.([]int)
	end, _ := this.end.Value.([]int)
	this.board[end[0]][end[1]] = 0

	switch direction {
	case "U":
		if head[0]-1 < 0 || this.board[head[0]-1][head[1]] == 1 {
			return -1
		}
		this.body.PushFront([]int{head[0] - 1, head[1]})
		this.board[head[0]-1][head[1]] = 1
	case "D":
		if head[0]+1 >= this.h || this.board[head[0]+1][head[1]] == 1 {
			return -1
		}
		this.body.PushFront([]int{head[0] + 1, head[1]})
		this.board[head[0]+1][head[1]] = 1
	case "R":
		if head[1]+1 >= this.w || this.board[head[0]][head[1]+1] == 1 {
			return -1
		}
		this.body.PushFront([]int{head[0], head[1] + 1})
		this.board[head[0]][head[1]+1] = 1
	case "L":
		if head[1]-1 < 0 || this.board[head[0]][head[1]-1] == 1 {
			return -1
		}
		this.body.PushFront([]int{head[0], head[1] - 1})
		this.board[head[0]][head[1]-1] = 1
	}

	head, _ = this.body.Front().Value.([]int)

	if this.foodIdx < len(this.food) && this.food[this.foodIdx][0] == head[0] && this.food[this.foodIdx][1] == head[1] {
		this.foodIdx++
		this.board[end[0]][end[1]] = 1
	} else {
		tmp := this.end.Prev()
		this.body.Remove(this.end)
		this.end = tmp
	}

	return this.body.Len() - 1
}
