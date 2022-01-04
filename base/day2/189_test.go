package day2

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	t.Run("#rotate", func(t *testing.T) {
		//got := rotate([]int{-1, -100, 3, 99}, 2)
		//fmt.Println("result:", got)
		rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	})
}
