package singlefizzbuzz

import (
	"fmt"
)

/*
1. Define a function called SingleFizzBuzz that have this behaviour:
1.1 It will receive a integer number n. By default, it return the integer number n without any operation
1.2 If n is divisible by 3, return Fizz.
1.3 If n is divisible by 5, return Buzz.
1.4 If n is divisible by 3 and 5, return FizzBuzz.
*/
func SingleFizzBuzz(n int64) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}
