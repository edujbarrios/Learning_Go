package main

import (
	"fmt"
)

const maze = `
+-+-+-+-+
S| |     |
+ + +-+-+ +
  |     | |
+ +-+ +-+ +
  |     | |
+ + +-+ +-+
  |   |   |
+ +-+ + +-+
  | E |   |
+-+-+-+-+
`

type point struct {
	x, y int
}

func (p point) add(r point) point {
	return point{p.x + r.x, p.y + r.y}
}

func (p point) at(grid [][]byte) byte {
	return grid[p.y][p.x]
}

func (p point) valid(grid [][]byte) bool {
	return p.x >= 0 && p.x < len(grid[0]) && p.y >= 0 && p.y < len(grid) && grid[p.y][p.x] != '|'
}

var dirs = [4]point{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func walk(grid [][]byte, start, end point) [][]byte {
	queue := []point{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			if next.valid(grid) {
				if grid[next.y][next.x] == ' ' {
					grid[next.y][next.x] = '*'
					queue = append(queue, next)
				}
			}
		}
	}

	return grid
}

func main() {
	grid := [][]byte{}
	for _, row := range strings.Split(maze, "\n") {
		grid = append(grid, []byte(row))
	}

	start, end := point{}, point{}
	for y, row := range grid {
		for x, cell := range row {
			switch cell {
			case 'S':
				start = point{x, y}
			case 'E':
				end = point{x, y}
			}
		}
	}

	grid = walk(grid, start, end)

	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}
