package main

import (
	"fmt"
	"os"
)

func ReadMaze(filename string) [][]int {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Printf("%s, %s, %s", pathError.Path, pathError.Op, pathError.Err)
		} else {
			panic(pathError)
		}
	}

	var row, col int

	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)

	for i, _ := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type Point struct {
	x, y int
}

var directions = [4]Point{
	{-1, 0}, {0, -1}, {0, 1}, {1, 0},
}

func (p Point) Move(r Point) Point {
	return Point{p.x + r.x, p.y + r.y}
}

func (p Point) At(grid [][]int) (int, bool) {
	if p.x >= 0 && p.x < len(grid) && p.y >= 0 && p.y < len(grid[0]) {
		return grid[p.x][p.y], true
	}
	return 0, false
}

func Walk(maze [][]int, start, end Point) [][]int {
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	//steps[start.x][start.y] = 1

	Q := []Point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range directions {
			next := cur.Move(dir)

			value, ok := next.At(maze)

			if !ok || value == 1 {
				continue
			}

			value, ok = next.At(steps)

			if !ok || value != 0 {
				continue
			}

			if next == start {
				continue
			}

			curStepValue, _ := cur.At(steps)
			steps[next.x][next.y] = curStepValue + 1
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := ReadMaze("./maze.in")

	steps := Walk(maze, Point{0, 0}, Point{len(maze) - 1, len(maze[0]) - 2})

	for _, row := range steps {
		for _, col := range row {
			fmt.Printf("%3d ", col)
		}
		fmt.Println()
	}
}
