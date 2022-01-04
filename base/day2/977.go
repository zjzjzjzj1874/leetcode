package day2

import "sort"

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//示例 1：
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]

// 很bug的常规解法 不如下面的解法
func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}

	sort.Ints(nums)
	return nums
}

// O(n)复杂度:双指针解法 ==> 推荐
func sortedSquaresWithOn(nums []int) []int {
	l := len(nums)
	i, j, r := 0, l, l-1
	res := make([]int, l)

	for i < j {
		if nums[i]*nums[i] >= nums[j]*nums[j] {
			res[r] = nums[i] * nums[i]
			i++
		} else {
			res[r] = nums[j] * nums[j]
			j--
		}
		r--
	}

	return res
}
