package day4

import (
	"testing"
)

func Test_reverseString(t *testing.T) {
	t.Run("#reverseString", func(t *testing.T) {
		reverseString([]byte("hello"))
	})
}
