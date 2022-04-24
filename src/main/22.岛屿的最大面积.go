package main

//695
//深度优先搜索
type pair22 struct {
	x, y int
}

var dirs = []pair22{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func maxAreaOfIsland(grid [][]int) int {
	var ans = 0
	for i, row := range grid {
		for j, _ := range row {
			ans = max(ans, getCnt(grid, i, j))
		}
	}
	return ans
}

func getCnt(grid [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != 1 {
		return 0
	}
	grid[x][y] = 0
	var res = 1
	for _, dir := range dirs {
		res += getCnt(grid, x+dir.x, y+dir.y)
	}
	return res
}

func main() {

}
