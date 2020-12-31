package _295_Find_Numbers_with_Even_Number_of_Digits

//Constraints make this question a lot easier
//Constraints:
//
//1 <= nums.length <= 500
//1 <= nums[i] <= 10^5
func findNumbers(nums []int) int {
	count :=0

	for i := 0 ; i < len(nums); i++ {

		if (nums[i]>9 && nums[i]<100) || (nums[i]>999 && nums[i]<10000) || nums[i] == 100000 {
			count++
		}
	}

	return count;
}