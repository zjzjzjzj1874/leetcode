package day2

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	t.Run("#searchInsert", func(t *testing.T) {
		got := sortedSquaresWithOn([]int{-9, -4, 0, 1, 3, 4})
		fmt.Println("result:", got)

	})
}
