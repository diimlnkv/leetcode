// Time:
// Space:
package main

func twoSum(nums []int, target int) []int {
	nmp := map[int]int{}
	for i, v := range nums {
		if _, ok := nmp[target-v]; ok {
			return []int{nmp[v], i}
		} else {
			nmp[v] = i
		}
	}
	return []int{}
}
