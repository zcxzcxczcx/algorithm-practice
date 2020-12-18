package main

func main() {

}
func intersect(nums1 []int, nums2 []int) []int {
	unionMap1 := map[int]int{}
	for _, v := range nums1 {
		unionMap1[v]++
	}
	unionMap2 := map[int]int{}
	for _, v := range nums2 {
		unionMap2[v]++
	}
	var result []int
	for k, v := range unionMap2 {
		if v > 0 && unionMap1[k] > 0 {
			count := min(v, unionMap1[k])
			for count > 0 {
				result = append(result, k)
				count--
			}

		}

	}
	return result

}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func intersect2(nums1 []int, nums2 []int) []int {
	unionMap1 := map[int]int{}
	for _, v := range nums1 {
		unionMap1[v]++
	}

	var result []int
	for _, v := range nums2 {
		if unionMap1[v] > 0 {
			result = append(result, v)
			unionMap1[v]--
		}
	}
	return result

}
