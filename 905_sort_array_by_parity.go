package LeetCodeGolangTest

func SortArrayByParity(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			t := nums[i]
			nums = append([]int{t}, append(nums[0:i], nums[i+1:]...)...)
		}
	}
	return nums
}
