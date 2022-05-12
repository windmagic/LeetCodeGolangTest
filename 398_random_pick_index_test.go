package LeetCodeGolangTest

import (
	"fmt"
	"testing"
)

func TestRandomPickIndex(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 3, 4, 4}
	target := 2
	obj := Constructor(nums)
	param1 := obj.Pick(target)
	fmt.Println(param1)
}
