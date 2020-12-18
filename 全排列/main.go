package main

import "fmt"

func main() {
	var s []int
	tem(&s)
	fmt.Println("ddddddddddddddd=%v", s)
}

func permute(nums []int) [][]int {
	numsLen := len(nums)
	flag := []int{}
	for i := 0; i < numsLen; i++ {
		flag = append(flag, 0)
	}
	temArr := []int{}
	res := [][]int{}
	backTrack(&flag, temArr, nums, &res)
	return res
}
func backTrack(flag *[]int, temArr []int, nums []int, res *[][]int) {
	if len(temArr) == len(nums) {

		*res = append(*res, temArr)
		fmt.Printf("resres,v=%v\n", *res)
		return
	}
	for i, v := range nums {
		fmt.Printf("dddddddddi=%v,v=%v\n", i, v)
		if (*flag)[i] != 1 {
			fmt.Printf("flagflag=%v\n", (*flag)[i])
			(*flag)[i] = 1
			temArr = append(temArr, v)
			fmt.Printf("	*temArr=%v\n", temArr)
			backTrack(flag, temArr, nums, res)
			(*flag)[i] = 0
			(temArr) = (temArr)[:len(temArr)-1]
			fmt.Printf("	* (*temArr) =(*temArr) [:len(*temArr)-1]=%v\n", temArr)
		}
	}
}
