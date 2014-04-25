// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "strings"

var letters string = "acdegilmnoprstuw"

func rehash(hash int64) string {
	result := ""
	for hash != 7 {
		mod := hash % 37
		hash = hash / 37
		result = letters[mod:mod+1] + result
	}
	return result
}

func hash(s string) int64 {
	var h int64 = 7
	for i := 0; i < len(s); i++ {
		k := strings.Index(letters, s[i:i+1])
		h = (h*37 + int64(k))
	}
	return h
}

func equal(test string, actual, expected int64) bool {
	if actual == expected {
		fmt.Printf("Test %v PASSED!\n", test)
		return true
	} else {
		fmt.Printf("Test %v failed. Actual: %v, Expected: %v\n", test, actual, expected)
		return false
	}
}

func main() {
	t1 := rehash(680131659347)
	t2 := rehash(910897038977002)
	fmt.Printf("Test = %v\nAnswer = %v\n", t1, t2)
	equal("leepadg", hash("leepadg"), 680131659347)
	equal("asparagus", hash("asparagus"), 910897038977002)
}
