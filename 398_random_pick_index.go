package LeetCodeGolangTest

import (
	"math/rand"
	"time"
)

type Solution map[int][]int

func Constructor(nums []int) Solution {
	end := Solution(map[int][]int{})

	for i := 0; i < len(nums); i++ {
		end[nums[i]] = append(end[nums[i]], i)
	}
	return end
}

func (this *Solution) Pick(target int) int {
	_len := len((*this)[target])
	rand.Seed(time.Now().Unix())
	index := rand.Intn(_len)
	return (*this)[target][index]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
