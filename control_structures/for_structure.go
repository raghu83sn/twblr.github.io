package control_structures

func collatzChainLength(n int) int {
	var count int
	count = 1
	recCollatz(n, &count)
	return count
}

func recCollatz(n int, count *int) int {
	if n == 1 {
		return 1
	} else if n%2 == 0 {
		(*count)++
		return recCollatz(n/2, count)
	} else if n%2 == 1 {
		(*count)++
		return recCollatz(3*n+1, count)
	} else {
		return n
	}

}
