package _26_Beautiful_Arrangement

func countArrangement(n int) int {
	count := 0
	visited := make([]bool, n + 1)

	calculate(n, 1, &count, visited)

	return count
}

func calculate( n int , pos int , count *int, visited []bool) {
	if pos > n {
		*count++
	}

	for i := 1 ; i <= n; i++ {

		if !visited[i] && (pos % i == 0 || i % pos == 0) {
			visited[i] = true
			calculate(n, pos + 1, count, visited)
			visited[i] = false
		}
	}
}
