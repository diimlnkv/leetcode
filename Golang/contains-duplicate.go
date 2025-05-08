// Time:
// Space:
package main

func containsDuplicate(nums []int) bool {
	var nm = map[int]int{}
	for _, val := range nums {
		if _, ok := nm[val]; ok {
			return true
		}
		nm[val] = 1
	}
	return false
}
