package day3

import (
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	t.Run("#moveZeroe", func(t *testing.T) {
		//moveZeroes([]int{0, 1, 3, 0, 5, 3})
		moveZeroesWithTp([]int{0, 1, 3, 0, 5, 3})
	})
}
