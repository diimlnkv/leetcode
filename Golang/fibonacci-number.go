// Time:
// Space:
package main

func fib(n int) int {
	if n < 2 {
		return n
	}

	pd := make([]int, n+1)
	pd[0], pd[1] = 0, 1

	for i := 2; i < n+1; i++ {
		pd[i] = pd[i-1] + pd[i-2]
	}

	return pd[n]
}
