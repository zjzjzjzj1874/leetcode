package day3

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	t.Run("#twoSum", func(t *testing.T) {
		fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	})
}
