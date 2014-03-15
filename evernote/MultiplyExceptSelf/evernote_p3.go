// https://evernote.com/careers/challenge.php
// Multiply Except Self

package main

import (
	"fmt"
)

var n int
var num []int64

func run() {
	var p int64 // product of all non-zeros in the set
	var zc int  // count number of zeros in the set

	fmt.Scanf("%d\n", &n)

	num = make([]int64, n)
	p = 1
	for i := 0; i < n; i++ {
		fmt.Scanf("%d\n", &num[i])
		if num[i] == 0 {
			if zc = zc + 1; zc > 1 {
				break
			}
		} else {
			p *= num[i]
		}
	}

	if zc > 1 { // more than one zeros, result is all zeros
		for i := 0; i < n; i++ {
			fmt.Printf("0\n")
		}
		return
	}

	if zc == 1 { // one zero
		for i := 0; i < n; i++ {
			if num[i] == 0 {
				fmt.Printf("%v\n", p)
			} else {
				fmt.Printf("0\n")
			}
		}
		return
	}

	// no zero
	for i := 0; i < n; i++ {
		fmt.Printf("%v\n", p/num[i])
	}
}

func main() {
	run()
}
