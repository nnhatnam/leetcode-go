package _689_Partitioning_Into_Minimum_Number_Of_Deci_Binary_Numbers

func minPartitions(n string) int {
	var max int32
	for _, v := range n {
		if v - '0' > max {
			max = v - '0'
		}
		if max == 9 {
			return 9
		}
	}
	return int(max)
}