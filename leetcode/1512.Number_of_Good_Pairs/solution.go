package _512_Number_of_Good_Pairs

func numIdenticalPairs(nums []int) int {
	var hashmap = map[int]int{}
	for _, v := range nums {
		if _, ok := hashmap[v]; ok {
			//the second time
			hashmap[v] += 1
		} else {
			hashmap[v] = 0
		}
	}

	result := 0
	for _, v := range hashmap {
		result += v*(v+1)/2

	}
	return result
}
