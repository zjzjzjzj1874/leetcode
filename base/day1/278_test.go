package day1

import (
	"fmt"
	"testing"
)

func Test_firstBad(t *testing.T) {
	t.Run("#第一个坏版本", func(t *testing.T) {
		fmt.Println(1 >> 1)
		fmt.Println(8 >> 2)

		bad = 4
		if got := firstBadVersion(5); got != bad {
			fmt.Println("查找失败:", got)
		} else {
			fmt.Println("查找成功:", got)
		}

		bad = 1
		if got := firstBadVersion(3); got != bad {
			fmt.Println("查找失败:", got)
		} else {
			fmt.Println("查找成功:", got)
		}

	})
}
