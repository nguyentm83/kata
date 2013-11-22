//http://codekata.pragprog.com/2007/01/kata_two_karate.html

package main

import (
	"log"
)

func chop(needle int, haystack []int) int {
	l := len(haystack)
	if l <= 0 {
		return -1
	}

	if l == 1 {
		if needle == haystack[0] {
			return 0
		} else {
			return -1
		}
	}

	// we know l > 1
	left := 0
	right := len(haystack) - 1

	for left <= right {
		mid := (left + right) >> 1
		k := haystack[mid]
		if needle == k {
			return mid
		}
		if needle < k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	} // end while

	return -1
}

func assert(t bool, note string) {
	if !t {
		log.Fatalf("%s FAILED !", note)
	}
}

func chop_test() {
	assert(-1 == chop(3, []int{}), "test 1")
	assert(-1 == chop(3, []int{1}), "test 2")
	assert(0 == chop(1, []int{1}), "test 3")

	assert(0 == chop(1, []int{1, 3, 5}), "test 4")
	assert(1 == chop(3, []int{1, 3, 5}), "test 5")
	assert(2 == chop(5, []int{1, 3, 5}), "test 6")
	assert(-1 == chop(0, []int{1, 3, 5}), "test 7")
	assert(-1 == chop(2, []int{1, 3, 5}), "test 8")
	assert(-1 == chop(4, []int{1, 3, 5}), "test 9")
	assert(-1 == chop(6, []int{1, 3, 5}), "test 10")

	assert(0 == chop(1, []int{1, 3, 5, 7}), "test 11")
	assert(1 == chop(3, []int{1, 3, 5, 7}), "test 12")
	assert(2 == chop(5, []int{1, 3, 5, 7}), "test 13")
	assert(3 == chop(7, []int{1, 3, 5, 7}), "test 14")
	assert(-1 == chop(0, []int{1, 3, 5, 7}), "test 15")
	assert(-1 == chop(2, []int{1, 3, 5, 7}), "test 17")
	assert(-1 == chop(4, []int{1, 3, 5, 7}), "test 18")
	assert(-1 == chop(6, []int{1, 3, 5, 7}), "test 19")
	assert(-1 == chop(8, []int{1, 3, 5, 7}), "test 20")

	log.Printf("ALL TEST PASSED")
}

func main() {
	chop_test()

}
