package main

import (
	"fmt"
	"os"
)

var dirs = [4]point{
	//上，左，右，下
	{0, -1},
	{-1, 0},
	{1, 0},
	{0, 1},
}

func readMaze(filename string) [][]int {
	reader, err := os.Open(filename)
	defer reader.Close()
	if err != nil {
		panic(err.Error())
	}
	var row, col int
	fmt.Fscanf(reader, "%d %d\r\n", &row, &col)
	maze := createGrid(row, col)
	for i := range maze {
		for j := range maze[i] {
			fmt.Fscanf(reader, "%d", &maze[i][j])
		}
		//读取需要换行,否则会读进0
		fmt.Fscanf(reader, "\r\n")
	}
	return maze
}

//创建二位数组
func createGrid(row, col int) [][]int {
	grid := make([][]int, row)
	for i, _ := range grid {
		grid[i] = make([]int, col)
	}
	return grid
}

func printGrid(grid [][]int) {
	cols := len(grid[0])
	fmt.Printf("   ")
	for i := 0; i < cols; i++ {
		fmt.Printf("%3d", i)
	}
	fmt.Println("\n-------------------------------")
	for rowNo, row := range grid {
		fmt.Printf("%2d|", rowNo)
		for _, col := range row {
			fmt.Printf("%3d", col)
		}
		fmt.Println()
	}
}

//走迷宫
func walk(maze [][]int, start point, end point) [][]int {
	steps := createGrid(len(maze), len(maze[0]))
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		//广度优先搜索开始
		for _, dir := range dirs {
			next := cur.add(dir)
			//迷宫在下一个格子的情况:
			at, ok := next.at(maze)
			//不能越界,这个点为0
			if !ok || at != 0 {
				continue
			}
			//没有走过
			step, ok := next.at(steps)
			if step > 0 {
				continue
			}
			//不能是起点
			if next == start {
				continue
			}
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			//将符合的点放到下一次搜寻的点队列里
			Q = append(Q, next)
		}
	}
	return steps
}

//画出路线，返回步数和路线的点数组
func drawRoute(steps [][]int, start, end point) (int, []point) {
	length, _ := end.at(steps)
	//初始化路径点数组
	route := make([]point, length+1)
	route[length] = end
	//遍历，寻找上一个点

	for target := length - 1; target >= 0; target = target - 1 {
		cur := route[target+1]
		//四个方向找目标点
		for _, dir := range dirs {
			last := cur.add(dir)
			at, b := last.at(steps)
			if !b || at != target {
				continue
			}
			route[target] = last
		}
	}
	route[0] = start
	return length, route
}

func main() {
	maze := readMaze("maze/maze.in")
	fmt.Println("Loading map...")
	printGrid(maze)
	fmt.Println("searching...")
	var start = point{0, 0}
	var end = point{len(maze) - 1, len(maze[0]) - 1}
	steps := walk(maze, start, end)
	printGrid(steps)
	stepCnt, route := drawRoute(steps, start, end)
	fmt.Printf("walk out maze need %d steps, route is %v", stepCnt, route)
}
