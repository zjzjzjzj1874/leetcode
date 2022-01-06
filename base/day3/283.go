package day3

import "fmt"

func moveZeroes(nums []int) {
	l := len(nums)
	res := make([]int, l)
	zeroNum := 0
	for i, num := range nums {
		if num == 0 {
			zeroNum++
			res[l-zeroNum] = 0
		} else {
			res[i-zeroNum] = num
		}
	}
	copy(nums, res)
	fmt.Println(nums)
}

// 双指针移动
func moveZeroesWithTp(nums []int) {
	idx, count := 0, 0 // count为0的数量,idx为非零元素移动偏移量
	for _, num := range nums {
		if num == 0 {
			count++
		} else {
			nums[idx] = num
			idx++
		}
	}

	for i := 0; i < count; i++ {
		nums[idx] = 0
		idx++
	}

}
