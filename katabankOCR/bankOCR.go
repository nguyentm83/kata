package main

import (
	"fmt"
	"log"
)

func assert(t bool, note string) {
	if !t {
		log.Fatalf("%s FAILED !", note)
	}
}

type bigdigit [3]string

/* GLOBAL VARIABLES */

var zero = bigdigit{
	" _ ",
	"| |",
	"|_|",
}

var one = bigdigit{
	"   ",
	"  |",
	"  |",
}

var two = bigdigit{
	" _ ",
	" _|",
	"|_ ",
}

var three = bigdigit{
	" _ ",
	" _|",
	" _|",
}

var four = bigdigit{
	"   ",
	"|_|",
	"  |",
}

var five = bigdigit{
	" _ ",
	"|_ ",
	" _|",
}

var six = bigdigit{
	" _ ",
	"|_ ",
	"|_|",
}

var seven = bigdigit{
	" _ ",
	"  |",
	"  |",
}

var eight = bigdigit{
	" _ ",
	"|_|",
	"|_|",
}

var nine = bigdigit{
	" _ ",
	"|_|",
	" _|",
}

var digitlist = []bigdigit{
	zero, one, two, three, four, five, six, seven, eight, nine,
}

func init() {

}

func equal(x, y bigdigit) bool {
	if x[0] == y[0] && x[1] == y[1] && x[2] == y[2] {
		return true
	}
	return false
}

func scan(d bigdigit) int {
	for i := 0; i < len(digitlist); i++ {
		if equal(d, digitlist[i]) {
			return i
		}
	}

	return -1
}

// user story 1 : take the account number picture from fax and convert it to account number
func OCR(d []bigdigit) string {
	result := ""
	for _, c := range d {
		v := scan(c)

		// assumption : v >= 0
		result += fmt.Sprintf("%d", v)
	}

	return result
}

func stringToDigit(s []string) []bigdigit {
	var digits []bigdigit

	for c := 0; c < 27; c += 3 {
		var d bigdigit
		d[0], d[1], d[2] = s[0][c:c+3], s[1][c:c+3], s[2][c:c+3]
		digits = append(digits, d)
	}

	return digits
}

func testOcr() {
	log.Printf("TEST OCR (USER STORY 1) STARTED ...")
	acct0 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		"| || || || || || || || || |",
		"|_||_||_||_||_||_||_||_||_|",
	}

	acct1 := []string{
		"                           ",
		"  |  |  |  |  |  |  |  |  |",
		"  |  |  |  |  |  |  |  |  |",
	}

	acct2 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		" _| _| _| _| _| _| _| _| _|",
		"|_ |_ |_ |_ |_ |_ |_ |_ |_ ",
	}

	acct3 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		" _| _| _| _| _| _| _| _| _|",
		" _| _| _| _| _| _| _| _| _|",
	}

	acct4 := []string{
		"                           ",
		"|_||_||_||_||_||_||_||_||_|",
		"  |  |  |  |  |  |  |  |  |",
	}

	acct5 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		"|_ |_ |_ |_ |_ |_ |_ |_ |_ ",
		" _| _| _| _| _| _| _| _| _|",
	}

	acct6 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		"|_ |_ |_ |_ |_ |_ |_ |_ |_ ",
		"|_||_||_||_||_||_||_||_||_|",
	}

	acct7 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		"  |  |  |  |  |  |  |  |  |",
		"  |  |  |  |  |  |  |  |  |",
	}

	acct8 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		"|_||_||_||_||_||_||_||_||_|",
		"|_||_||_||_||_||_||_||_||_|",
	}

	acct9 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		"|_||_||_||_||_||_||_||_||_|",
		" _| _| _| _| _| _| _| _| _|",
	}

	acct10 := []string{
		"    _  _     _  _  _  _  _ ",
		"  | _| _||_||_ |_   ||_||_|",
		"  ||_  _|  | _||_|  ||_| _|",
	}

	assert("000000000" == OCR(stringToDigit(acct0)), "test 0")
	assert("111111111" == OCR(stringToDigit(acct1)), "test 1")
	assert("222222222" == OCR(stringToDigit(acct2)), "test 2")
	assert("333333333" == OCR(stringToDigit(acct3)), "test 3")
	assert("444444444" == OCR(stringToDigit(acct4)), "test 4")
	assert("555555555" == OCR(stringToDigit(acct5)), "test 5")
	assert("666666666" == OCR(stringToDigit(acct6)), "test 6")
	assert("777777777" == OCR(stringToDigit(acct7)), "test 7")
	assert("888888888" == OCR(stringToDigit(acct8)), "test 8")
	assert("999999999" == OCR(stringToDigit(acct9)), "test 9")
	assert("123456789" == OCR(stringToDigit(acct10)), "test 10")

}

func main() {
	testOcr()
}
