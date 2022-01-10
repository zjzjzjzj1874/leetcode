package day3

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//说明:
//
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。
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
