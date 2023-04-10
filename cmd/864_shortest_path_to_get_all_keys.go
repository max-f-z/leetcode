package main

import (
	"bytes"
	"encoding/gob"
	"math"
)

type pathInfo struct {
	keysHave map[byte]bool
	x        int
	y        int
	moves    int
	// visited  [][]map[int]bool
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func calculateKey(m map[byte]bool) int {
	res := 0
	for k := range m {
		res += int(math.Pow(2, float64(k-'a')))
	}
	return res
}

//lint:ignore U1000 unused
func shortestPathAllKeys(grid []string) int {
	h := len(grid)
	w := len(grid[0])
	var startX, startY int

	keys := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] >= 'a' && grid[i][j] <= 'f' {
				keys++
			}

			if grid[i][j] == '@' {
				startX = i
				startY = j
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	p := pathInfo{
		x:        startX,
		y:        startY,
		moves:    0,
		keysHave: map[byte]bool{},
	}
	visited := make([][]map[int]bool, h)
	for i := 0; i < h; i++ {
		visited[i] = make([]map[int]bool, w)
		for j := 0; j < w; j++ {
			visited[i][j] = make(map[int]bool)
		}
	}

	visited[startX][startY][calculateKey(p.keysHave)] = true

	queue := []pathInfo{p}
	for len(queue) > 0 {
		newQ := []pathInfo{}

		for _, p := range queue {

			for _, dir := range directions {
				x := p.x + dir[0]
				y := p.y + dir[1]

				if x >= 0 && y >= 0 && x < h && y < w {
					if grid[x][y] == '#' {
						continue
					}

					if visited[x][y][calculateKey(p.keysHave)] {
						continue
					}

					if grid[x][y] == '.' || grid[x][y] == '@' {
						pp := pathInfo{
							x:     x,
							y:     y,
							moves: p.moves + 1,
						}

						ppKeysHave := make(map[byte]bool)
						deepCopy(&ppKeysHave, p.keysHave)
						pp.keysHave = ppKeysHave

						visited[x][y][calculateKey(pp.keysHave)] = true

						newQ = append(newQ, pp)
					}

					if grid[x][y] >= 'a' && grid[x][y] <= 'f' {
						pp := pathInfo{
							x:     x,
							y:     y,
							moves: p.moves + 1,
						}

						ppKeysHave := make(map[byte]bool)
						deepCopy(&ppKeysHave, p.keysHave)
						pp.keysHave = ppKeysHave

						pp.keysHave[grid[x][y]] = true

						visited[x][y][calculateKey(pp.keysHave)] = true

						if len(pp.keysHave) == keys {
							return pp.moves
						}

						newQ = append(newQ, pp)
					}

					if grid[x][y] >= 'A' && grid[x][y] <= 'F' {
						if p.keysHave[grid[x][y]+32] {
							pp := pathInfo{
								x:     x,
								y:     y,
								moves: p.moves + 1,
							}

							ppKeysHave := make(map[byte]bool)
							deepCopy(&ppKeysHave, p.keysHave)
							pp.keysHave = ppKeysHave

							visited[x][y][calculateKey(pp.keysHave)] = true

							newQ = append(newQ, pp)
						}
					}
				}
			}
		}
		queue = newQ
	}

	return -1
}
