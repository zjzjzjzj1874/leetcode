package day2

//给你一个数组，将数组中的元素向右轮转 k个位置，其中k是非负数。
/*
示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
*/
/*
示例2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
*/
/*
示例3:
输入：nums = [-1,-100,3], k = 4
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [3,-1,-100]
向右轮转 2 步: [-100,3,-1]
向右轮转 3 步: [-1,-100,3]
向右轮转 4 步: [3,-1,-100]
*/
func rotate(nums []int, k int) {
	l := len(nums)
	res := make([]int, l)
	for i := range nums {
		res[(i+k)%l] = nums[i]
	}
	copy(nums, res)
}
