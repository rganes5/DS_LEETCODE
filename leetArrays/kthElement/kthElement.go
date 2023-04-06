package main

import "fmt"

func kthElement(userSlice []int, kthLargest int, kthSmallest int) []int {
	//SORTING
	for i := 0; i < len(userSlice)-1; i++ {
		for j := i + 1; j < len(userSlice); j++ {
			if userSlice[i] > userSlice[j] {
				temp := userSlice[i]
				userSlice[i] = userSlice[j]
				userSlice[j] = temp
			}
		}
	}
	fmt.Printf("\nThe %v largest element is %v", kthLargest, userSlice[len(userSlice)-kthLargest])
	fmt.Printf("\nThe %v smallest element is %v", kthSmallest, userSlice[kthSmallest-1])
	return userSlice

}

func main() {
	var kthLargest, kthSmallest int
	userSlice := []int{1, 9, 2, 8, 4, 7, 6, 3}
	fmt.Println("Enter the kth smallest number that you would like to check ")
	fmt.Scan(&kthSmallest)
	fmt.Println("Enter the kth largest number that you would like to check ")
	fmt.Scan(&kthLargest)
	kth := kthElement(userSlice, kthLargest, kthSmallest)
	fmt.Printf("\n%v", kth)
}
