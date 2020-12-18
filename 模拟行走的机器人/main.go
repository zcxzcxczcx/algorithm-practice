package main

func main() {

}
func robotSim(commands []int, obstacles [][]int) int {
	obstaclesMap := map[int][]int{}
	for _, v := range obstacles {
		obstaclesMap[v[0]] = append(obstaclesMap[v[0]], v[1])
	}
	x := 0
	y := 0
	direct := 0 //0:上、1:右、2:下、3:左
	store := [][]int{}
	for _, v := range commands {
		if v >= 1 && v <= 9 {
			tempx := x
			tempy := y
			for i := 1; i <= v; i++ {
				if direct == 0 {
					tempy = tempy + 1

				} else if direct == 1 {
					tempx = tempx + 1
				} else if direct == 2 {
					tempy = tempy - 1
				} else {
					tempx = tempx - 1
				}
				flag := 0
				for _, v1 := range obstaclesMap[tempx] {
					if tempy == v1 {
						flag = 1
					}
				}
				if flag == 1 {
					break
				} else {
					x = tempx
					y = tempy
					store = append(store, []int{x, y})
				}
			}
		}
		if v == -1 {
			direct = (direct + 1) % 4
		}
		if v == -2 {
			direct = (direct + 3) % 4
		}

	}
	max := 0
	for _, v := range store {
		num := v[0]*v[0] + v[1]*v[1]
		if num > max {
			max = num
		}
	}
	return max

}
