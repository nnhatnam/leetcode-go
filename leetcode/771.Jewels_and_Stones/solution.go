package _71_Jewels_and_Stones

func numJewelsInStones(jewels string, stones string) int {
	jMap := map[rune]int{}
	for _, jewel := range jewels {
		jMap[jewel] = 1
	}
	sum := 0
	for _, stone := range stones {
		sum += int(jMap[stone])
	}
	return sum
}
