package _282_Group_the_People_Given_the_Group_Size_They_Belong_To


func groupThePeople(groupSizes []int) [][]int {
	hashmap := map[int][]int{}
	result := [][]int{}
	for person, groupSize := range groupSizes {
		hashmap[groupSize] = append(hashmap[groupSize], person)

		if len(hashmap[groupSize]) == groupSize {
			result = append(result, hashmap[groupSize])
			delete(hashmap, groupSize)
		}
	}

	return result

}
