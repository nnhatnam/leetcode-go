package _006_ZigZac_Conversion

//need to know zigzac pattern
//improve later

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	m := make(map[int]string)
	for i := 0; i < numRows; i++ {
		m[i] = ""
	}
	step := 1
	flag := true
	for _, char := range s {


		if flag {
			m[step - 1] +=  string(char)
		} else {
			m[numRows - step] += string(char)
		}

		if step % numRows == 0 {
			flag = !flag
			step = 2
		} else {
			step++
		}
	}
	result := ""
	for i := 0; i < numRows; i++ {
		result += m[i]
	}
	return result
}
