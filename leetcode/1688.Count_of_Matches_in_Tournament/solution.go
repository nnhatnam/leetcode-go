package _688_Count_of_Matches_in_Tournament

func numberOfMatches(n int) int {
	total := 0
	for teams := n; teams >= 2; {
		var nMatches int
		if teams % 2 == 0 {
			nMatches = teams / 2
			teams = teams / 2
		} else {
			nMatches = (teams  - 1) / 2
			teams = (teams - 1) / 2 + 1
		}
		total += nMatches
	}
	return total
}
