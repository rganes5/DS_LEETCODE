package main

import "fmt"

func sumOfUnique(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		temp := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				temp++
			}
		}
		if temp <= 1 {
			sum = sum + nums[i]
			fmt.Println(nums[i])
		}
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 3, 5}
	sum := sumOfUnique(nums)
	fmt.Printf("The sum is %v", sum)

}
