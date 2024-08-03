package main

import (
	"fmt"
)

// Write a function that accepts an array of 10 integers (between 0 and
// 9), that returns a string of those numbers in the form of a phone
// number.
//
// EXAMPLE:
// CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})
// returns "(123) 456-7890"
func CreatePhoneNumber(numbers [10]uint) (phoneNumber string) {
	var combine = func(nums []uint) (numstring string) {
		for _, n := range nums {
			numstring += fmt.Sprintf("%v", n)
		}
		return
	}

	phoneNumber = fmt.Sprintf(
		"(%s) %s-%s",
		combine(numbers[:3]),
		combine(numbers[3:6]),
		combine(numbers[6:]),
	)
	return
}
