package Check_Array_Formation_Through_Concatenation

func canFormArray(arr []int, pieces [][]int) bool {
	hashmap := map[int][]int{}
	for _, v := range pieces {
		hashmap[v[0]] = v
	}

	if _, ok := hashmap[arr[0]]; !ok {
		return false
	}

	for i := 0; i < len(arr); i++ {
		piece, ok := hashmap[arr[i]]
		if ok {
			k := 0
			for ; k < len(piece) - 1; k++ {
				if i + k >= len(arr) {
					return false
				} else if piece[k] != arr[i + k] {
					return false
				}
			}
			i = i + k

		} else {
			return false
		}
	}
	return true
}
