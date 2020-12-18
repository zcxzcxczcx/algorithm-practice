package main

import "fmt"

func main() {

	var a = [9]int{}
	a[0] = 0
	a[0] = 1
	for i := 2; i < 9; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	fmt.Printf("a[8]==========%v\n", a[8])
}
