package main

type point struct {
	i, j int
}

func (p point) add(dir point) point {
	return point{p.i + dir.i, p.j + dir.j}
}

//查看迷宫当前点的情况
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i > len(grid)-1 {
		return 0, false
	}
	if p.j < 0 || p.j > len(grid[0])-1 {
		return 0, false
	}
	return grid[p.i][p.j], true
}
