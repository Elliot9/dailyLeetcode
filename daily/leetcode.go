package Daily

import "fmt"

func FindTheTownJudge() {
	trust := make([][]int, 0)
	result := findJudge(2, trust)
	fmt.Println(result)
}

func findJudge(n int, trust [][]int) int {
	trustCount := make([]int, n+1)
	trustedCount := make([]int, n+1)
	for _, t := range trust {
		trustCount[t[0]]++
		trustedCount[t[1]]++
	}

	for i := 1; i <= n; i++ {
		if trustCount[i] == 0 && trustedCount[i] == n-1 {
			return i
		}
	}

	return -1
}
