package main

import "fmt"

func activity(userSlice []int) {
	max := userSlice[0]
	for i := 0; i < len(userSlice); i++ {
		if max < userSlice[i] {
			max = userSlice[i]
		}
		userSlice[i] += 3
	}

	for i := 0; i < len(userSlice); i++ {
		if max < userSlice[i] {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
	fmt.Println(max)
	fmt.Println(userSlice)
}

func main() {
	userSlice := []int{1, 5, 4, 11, 7, 3, 10, 6}
	activity(userSlice)

}
