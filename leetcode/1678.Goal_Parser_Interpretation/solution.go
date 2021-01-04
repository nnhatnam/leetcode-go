package _678_Goal_Parser_Interpretation

import "strings"

func interpret(command string) string {
	var s strings.Builder
	i, n :=  0, len(command)
	for i < n {
		if command[i:i+1] == "G" {
			s.WriteString("G")
			i++
		} else if command[i:i+2] == "()" {
			s.WriteString("o")
			i += 2
		} else {
			s.WriteString("al")
			i += 4
		}
	}
	return s.String()
}
