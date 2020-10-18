package _011_Container_With_Most_Water

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	lo := 0
	hi := len(height) - 1
	max := 0
	for lo < hi {
		minHeight := min(height[lo],  height[hi])
		area := (hi - lo) * minHeight
		if area > max {
			max = area
		}
		if height[lo] == minHeight {
			lo++
		} else {
			hi--
		}
	}
	return max
}
