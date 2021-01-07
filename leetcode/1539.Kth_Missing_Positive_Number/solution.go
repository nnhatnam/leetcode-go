package _539_Kth_Missing_Positive_Number

func findKthPositive(arr []int, k int) int {
	j := 0
	count := 0
	for i := 1 ; ; i++ {


		if j < len(arr) && arr[j] == i {
			j++
		} else {
			count++
		}

		if count == k {
			return i
		}


	}
	return 0
}
