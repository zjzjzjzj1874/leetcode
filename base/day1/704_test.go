package day1

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	t.Run("#704二分查找", func(t *testing.T) {
		if got := search([]int{0, 1, 3, 4, 7, 8, 13, 20}, 5); got != -1 {
			fmt.Println("查找失败:", got)
		} else {
			fmt.Println("查找成功:", got)
		}

		if got := search([]int{0, 1, 3, 5, 7, 8, 13, 20}, 5); got != 3 {
			fmt.Println("查找失败:", got)
		} else {
			fmt.Println("查找成功:", got)
		}
		if got := search([]int{-1, 0, 3, 5, 9, 12}, -1); got != 0 {
			fmt.Println("查找失败:", got)
		} else {
			fmt.Println("查找成功:", got)
		}
	})
}
