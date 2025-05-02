// Time:
// Space:
package main

func tribonacci(n int) int {
	if n < 3 {
		if n == 0 {
			return 0
		}
		return 1
	}
	pd := make([]int, n+1)
	pd[0], pd[1], pd[2] = 0, 1, 1

	for i := 3; i < n+1; i++ {
		pd[i] = pd[i-3] + pd[i-2] + pd[i-1]
	}
	return pd[n]
}
