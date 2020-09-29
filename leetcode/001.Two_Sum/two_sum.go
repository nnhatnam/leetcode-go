package twosum

//bruce force
//Time complexity: O(n^2)
//Space complexity: O(1)
func twoSum1(nums []int, target int) []int {
	for i, num := range nums{
		for k, v := range nums {
			if (k != i) && (num + v == target) {
				return []int{i, k}
			}
		}
	}
	return []int{}
}


//Hashing
////Time complexity: O(n)
//Space complexity: O(n)
func twoSum2(nums []int, target int) []int {
	hmap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		v, ok := hmap[num]
		if ok {
			return []int{v, i}
		}
		hmap[complement] = i
	}
	return []int{}
}

//This is not a solution
//Sorting and Binary Search if we want to return number instead of index
//Time complexity: O(nlogn)
//Space complexity: O(1)
//func twoSum3(nums []int, target int) []int {
//	sort.Ints(nums)
//	for i, num := range nums {
//
//		complement := target - num
//		siblingIndex := sort.Search(len(nums), func(j int) bool {
//			return nums[j] >= complement
//		})
//		if siblingIndex < len(nums) && nums[siblingIndex] == complement {
//			return []int{i, siblingIndex}
//		}
//	}
//	return []int{}
//}