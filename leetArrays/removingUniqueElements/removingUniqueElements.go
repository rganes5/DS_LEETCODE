package main

import "fmt"

func unique(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				nums = append(nums[:i], nums[j:]...)
				j--
			}
		}
	}
	fmt.Println(nums)
	return len(nums)
}
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// nums := []int{1, 2, 3, 4, 5}
	k := unique(nums)
	fmt.Printf("%v", k)
}
