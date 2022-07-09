/*
An "even ended number" is a number whose first and last digit are the same.

You mission, should you choose to accept it, is to count how many "even ended numbers" are
there that are a multiplication of two 4 digit numbers.
*/

package main

import (
	"fmt"
)

func main() {
	// n := 42
	// s := fmt.Sprintf("%d", n)
	// fmt.Printf("s = %q (type %T)\n", s, s)
	
	counter:= 0
	
	for i := 1000; i <= 9999; i++ {
		for j:= i; j <= 9999; j++ {
			num := i * j
			num_str := fmt.Sprintf("%d", num)
			if last_index:=len(num_str) - 1; num_str[0] == num_str[last_index] {
				counter++
			}
		}
	}

	fmt.Printf("Total Even Ended Numbers are : %d", counter);
}
