// Time:
// Space:
package main

func findDisappearedNumbers(nums []int) []int {
	sortedNums := make([]int, len(nums))
	var ans []int

	for _, v := range nums {
		sortedNums[v-1] = v
	}

	for i, v := range sortedNums {
		if v == 0 {
			ans = append(ans, i+1)
		}

	}
	return ans
}
