package day1

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	t.Run("#searchInsert", func(t *testing.T) {
		if got := searchInsert([]int{0, 1, 3, 4, 7, 8, 13, 20}, 5); got != 4 {
			fmt.Println("查找失败:", got)
		} else {
			fmt.Println("查找成功:", got)
		}

		if got := searchInsert([]int{0, 1, 3, 5, 7, 8, 13, 20}, 4); got != 3 {
			fmt.Println("查找失败:", got)
		} else {
			fmt.Println("查找成功:", got)
		}

		if got := searchInsert([]int{1, 3, 5, 6}, 7); got != 4 {
			fmt.Println("查找失败:", got)
		} else {
			fmt.Println("查找成功:", got)
		}
	})
}
