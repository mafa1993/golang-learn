package main

// 广度优先实现迷宫算法
import (
	"fmt"
	"os"
)

// 用来记录位置坐标
type point struct {
	row int
	col int
}

var dirs = [4]point{ // 定义遍历方向，用相对位置来计算
	{-1, 0}, // 上 ，行号-1  列号布标
	{0, -1}, // 左
	{1, 0},  // 下
	{0, 1},  // 右
}

func main() {
	// 获取迷宫,迷宫文件必须为lf  crlf会出错
	maze := readMaze("maze.in")

	step := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	fmt.Println("step为", step)
}

// 获取迷宫
func readMaze(filename string) [][]int {
	file, err := os.Open(filename) // os.Open只读打开，可以使用os.Openfile() 来指定读写
	if nil != err {
		panic(err)
	}
	var row, col int // 用来存储行列
	fmt.Fscanf(file, "%d %d", &row, &col)
	fmt.Println("行列", row, col)

	// 获取迷宫
	var maze [][]int = make([][]int, row) // 用来存储迷宫 slice

	for k := range maze { //遍历行
		maze[k] = make([]int, col) // 每一行的元素
		for i := range maze[k] {   // 遍历列

			fmt.Fscanf(file, "%d", &maze[k][i])
		}
	}

	fmt.Println("迷宫为", maze)

	return maze
}

// 走迷宫
/**
 *  @param maze 迷宫
 *  @param start  迷宫开始坐标  这里使用结构体类型来表示，也可以使用数组
 *  @param end  迷宫结束坐标
 */
func walk(maze [][]int, start, end point) [][]int {
	step := make([][]int, len(maze)) // 用来存储所有走过的点 以及这个点是第几步走的
	for i := range step {
		step[i] = make([]int, len(maze[i]))
	}
	queue := []point{start} // 存储将要探索的点，开始为start

	// 队列有数据才进行
	for len(queue) > 0 {
		cur := queue[0]   // 获取队列值
		queue = queue[1:] // 获取到的值移除

		// 到达重点 退出
		if cur == end {
			break
		}
		for _, v := range dirs {
			next := cur.add(v) // 下一位置为对队列的值进行上下左右移动

			// 对next 进行判断，1 maze中是0,迷宫中可以走的路  step为0，代表没走过  不能为负数  不能为start
			//fmt.Println("next", next)
			val, ok := next.val(maze) // val 为next这个点在迷宫的值

			if !ok || val == 1 { // 不能走的情况，越界和迷宫中的墙
				continue
			}

			val, ok = next.val(step)

			if !ok || val != 0 { // 不能走的情况 越界和已经走过的
				continue
			}

			if next == start { // 不能重复走起点
				continue
			}

			queue = append(queue, next)            // 将符合条件的点加入到队列 下次使用
			curStep, _ := cur.val(step)            // 当前位置 step的值
			step[next.row][next.col] = curStep + 1 // 下一步的step值 为当前所在位置+1
		}
	}

	return step
}

// 用于计算坐标
func (p point) add(r point) point {
	return point{p.row + r.row, p.col + r.col}
}

// 返回point位置在maze或者step对应位置的值
func (p point) val(grid [][]int) (int, bool) {
	if p.col < 0 || p.row < 0 { // 行列数不能小于0
		return 0, false
	}
	if p.col >= len(grid[0]) || p.row >= len(grid) { // 不能越迷宫边界
		return 0, false
	}

	return grid[p.row][p.col], true
}
