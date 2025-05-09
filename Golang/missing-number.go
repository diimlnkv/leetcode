// Time:
// Space:
package main

// XOR method
func missingNumber(nums []int) int {
	xor := 0
	n := len(nums) + 1

	for i := 0; i < n; i++ {
		xor ^= i
	}

	for _, i := range nums {
		xor ^= i
	}

	return xor
}
