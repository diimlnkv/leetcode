// Time:
// Space:
package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	var mp = map[int]int{}

	for _, val := range nums {
		mp[val] = mp[val] + 1
	}
	return topKValues(mp, k)
}

type kv struct {
	Key   int
	Value int
}

func topKValues(m map[int]int, k int) []int {
	var kvs = make([]kv, 0, len(m))

	for k, v := range m {
		kvs = append(kvs, kv{k, v})
	}

	sort.Slice(kvs, func(i, j int) bool { return kvs[i].Value > kvs[j].Value })

	if k > len(kvs) {
		k = len(kvs)
	}

	var kFrequent = make([]int, k)

	for i := 0; i < k; i++ {
		kFrequent[i] = kvs[i].Key
	}

	return kFrequent
}
