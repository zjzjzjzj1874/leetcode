package day1

import "sort"

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//请必须使用时间复杂度为 O(log n) 的算法。

func searchInsert(nums []int, target int) int {
	return sort.SearchInts(nums, target)

	i, j := 0, len(nums)
	for i < j {
		m := int((uint(i) + uint(j)) >> 1)
		if nums[m] == target {
			return m
		} else {
			if nums[m] > target {
				j = m
			} else {
				i = m + 1
			}
		}
	}
	return i
}
