// Calculate maximal value in a slice
package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	fmt.Println(nums)
	max_num := nums[0]
	for _, num := range nums[1:] {
		if num > max_num {
			max_num = num
		}
	}
	fmt.Printf("The maximal value in given slice is: %d", max_num)
}
