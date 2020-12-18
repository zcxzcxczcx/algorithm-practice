package main

func main() {

}
func longestCommonSubsequence(text1 string, text2 string) int {
	var arr [][]int
	lenText1 := len(text1)
	lenText2 := len(text2)
	for i := 0; i < lenText1; i++ {
		arr = append(arr, []int{})
		for j := 0; j < lenText2; j++ {
			arr[i] = append(arr[i], 0)
		}
	}
	flag1 := false
	for i := 0; i < lenText2; i++ {
		if text2[i] == text1[0] {
			flag1 = true
		}
		if flag1 {
			arr[0][i] = 1
		}
	}

	flag2 := false
	for i := 0; i < lenText1; i++ {
		if text1[i] == text2[0] {
			flag2 = true
		}
		if flag2 {
			arr[i][0] = 1
		}
	}
	for i := 1; i < lenText1; i++ {
		for j := 1; j < lenText2; j++ {

			if text1[i] == text2[j] {

				arr[i][j] = arr[i-1][j-1] + 1

			} else {

				arr[i][j] = max(arr[i-1][j], arr[i][j-1])

			}
		}

	}
	return arr[lenText1-1][lenText2-1]

}
func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
