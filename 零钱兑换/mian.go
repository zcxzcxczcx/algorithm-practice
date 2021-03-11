package main

import "math"

func main() {

}
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	var storeArr = make([]int, amount)
	return coinCount(coins, amount, storeArr)

}
func coinCount(coins []int, amount int, storeArr []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if storeArr[amount-1] != 0 {
		return storeArr[amount-1]
	}
	var min = math.MaxUint32
	for _, coin := range coins {
		var temp int
		temp = coinCount(coins, amount-coin, storeArr)
		if temp >= 0 && temp < min {
			min = temp + 1
		}

	}

	if min == math.MaxUint32 {
		min = -1
	}
	storeArr[amount-1] = min
	return min

}
