// http://codingdojo.org/cgi-bin/wiki.pl?KataBankOCR
package main

import (
	"fmt"
	"log"
	"strings"
)

func assert(t bool, note string) {
	if !t {
		log.Fatalf("%s FAILED !", note)
	}
}

const (
	VALID = 0
	ILL   = 1
	ERR   = 2
)

type bigdigit [3]string

type element struct {
	graphic       *bigdigit // graphic representation
	possibilities []int     // possible character
}

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

var alldigit []element

func isReplaceable(a, b *bigdigit) bool {
	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			//			if a[i][j] == ' ' && (b[i][j] == '|' || b[i][j] == '_') {
			// if a[i][j] != b[i][j] && (a[i][j] == ' ' || b[i][j] == ' ') {
			if a[i][j] != b[i][j] {
				count++
			} else if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	if count == 1 {
		return true
	}

	return false
}

func init() {
	alldigit = make([]element, 10)
	alldigit[0] = element{&zero, []int{}}
	alldigit[1] = element{&one, []int{}}
	alldigit[2] = element{&two, []int{}}
	alldigit[3] = element{&three, []int{}}
	alldigit[4] = element{&four, []int{}}
	alldigit[5] = element{&five, []int{}}
	alldigit[6] = element{&six, []int{}}
	alldigit[7] = element{&seven, []int{}}
	alldigit[8] = element{&eight, []int{}}
	alldigit[9] = element{&nine, []int{}}

	// initialize the possibilities
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j {
				continue
			}
			if k := isReplaceable(alldigit[i].graphic, alldigit[j].graphic); k == true {
				alldigit[i].possibilities = append(alldigit[i].possibilities, j)
			}
		}
	}

	/*
		for i, v := range alldigit {
			fmt.Printf("%d %v\n", i, v.possibilities)
		}
	*/
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
		if v := scan(c); v >= 0 {
			result += fmt.Sprintf("%d", v)
		} else {
			result += "?"
		}
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

func isValid(s string) bool {
	// calculate checksum
	d9 := int(s[0]) - '0'
	d8 := int(s[1]) - '0'
	d7 := int(s[2]) - '0'
	d6 := int(s[3]) - '0'
	d5 := int(s[4]) - '0'
	d4 := int(s[5]) - '0'
	d3 := int(s[6]) - '0'
	d2 := int(s[7]) - '0'
	d1 := int(s[8]) - '0'
	chksum := (d1 + 2*d2 + 3*d3 + 4*d4 + 5*d5 + 6*d6 + 7*d7 + 8*d8 + 9*d9) % 11
	return chksum == 0
}

/* TESTING */

func TestStory1() {
	log.Printf("TEST USER STORY 1 STARTED ...")
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

func validate(s string) int {
	if strings.Contains(s, "?") {
		return ILL
	}

	if isValid(s) {
		return VALID
	}

	return ERR
}

func validatestr(s string) string {
	if v := validate(s); v == ILL {
		return s + " ILL"
	} else if v == ERR {
		return s + " ERR"
	}
	return s
}

func TestStory3() {
	log.Printf("TEST USER STORY 3 STARTED ...")

	acct0 := []string{
		" _  _  _  _  _  _  _  _    ",
		"| || || || || || || ||_   |",
		"|_||_||_||_||_||_||_| _|  |",
	}

	acct1 := []string{
		"    _  _  _  _  _  _     _ ",
		"|_||_|| || ||_   |  |  | _ ",
		"  | _||_||_||_|  |  |  | _|",
	}

	acct2 := []string{
		"    _  _     _  _  _  _  _ ",
		"  | _| _||_| _ |_   ||_||_|",
		"  ||_  _|  | _||_|  ||_| _ ",
	}

	assert("000000051" == validatestr(OCR(stringToDigit(acct0))), "test 11")
	assert("49006771? ILL" == validatestr(OCR(stringToDigit(acct1))), "test 12")
	assert("1234?678? ILL" == validatestr(OCR(stringToDigit(acct2))), "test 13")
	assert("664371495 ERR" == validatestr("664371495"), "test 14")
	assert("86110??36 ILL" == validatestr("86110??36"), "test 15")
}

func fixAccount(s []string) string {
	acctNo := OCR(stringToDigit(s))
	result := []string{acctNo}

	if status := validate(acctNo); status == ERR {
		for i := 0; i < len(acctNo); i++ {
			keep := int(acctNo[i]) - '0'
			safe := acctNo
			for j := 0; j < len(alldigit[keep].possibilities); j++ {
				t := []byte(acctNo)
				t[i] = byte(alldigit[keep].possibilities[j]) + 48
				acctNo = string(t)

				if validate(acctNo) == VALID {
					result = append(result, acctNo)
				}
			}
			acctNo = safe
		}
	} else if status == ILL {
		//fmt.Println(acctNo)
		v := strings.Index(acctNo, "?")
		// invariant: when s == ILL, v >= 0

		d := stringToDigit(s)[v]
		//fmt.Println(d)
		for i := 0; i < 10; i++ {
			if isReplaceable(&d, alldigit[i].graphic) {
				t := []byte(acctNo)
				//	fmt.Println(t)
				t[v] = byte(i) + 48
				acctNo = string(t)
				//	fmt.Println(acctNo)

				if validate(acctNo) == VALID {
					result = append(result, acctNo)
					break
				}
			}
		}

	} else {
		return acctNo
	}

	return strings.Join(result[1:], ",")
}

func TestStory4() {
	log.Printf("TEST USER STORY 4 STARTED ...")

	tc0 := []string{
		"                           ",
		"  |  |  |  |  |  |  |  |  |",
		"  |  |  |  |  |  |  |  |  |",
	}
	tc1 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		"  |  |  |  |  |  |  |  |  |",
		"  |  |  |  |  |  |  |  |  |",
	}
	tc2 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		" _|| || || || || || || || |",
		"|_ |_||_||_||_||_||_||_||_|",
	}
	tc3 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		" _| _| _| _| _| _| _| _| _|",
		" _| _| _| _| _| _| _| _| _|",
	}
	tc4 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		"|_||_||_||_||_||_||_||_||_|",
		"|_||_||_||_||_||_||_||_||_|",
	}
	tc5 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		"|_ |_ |_ |_ |_ |_ |_ |_ |_ ",
		" _| _| _| _| _| _| _| _| _|",
	}
	//=> 555555555 AMB ['555655555', '559555555']
	tc6 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		"|_ |_ |_ |_ |_ |_ |_ |_ |_ ",
		"|_||_||_||_||_||_||_||_||_|",
	}
	//=> 666666666 AMB ['666566666', '686666666']
	tc7 := []string{
		" _  _  _  _  _  _  _  _  _ ",
		"|_||_||_||_||_||_||_||_||_|",
		" _| _| _| _| _| _| _| _| _|",
	}
	//=> 999999999 AMB ['899999999', '993999999', '999959999']
	tc8 := []string{
		"    _  _  _  _  _  _     _ ",
		"|_||_|| || ||_   |  |  ||_ ",
		"  | _||_||_||_|  |  |  | _|",
	}
	//=> 490067715 AMB ['490067115', '490067719', '490867715']

	tc9 := []string{
		"    _  _     _  _  _  _  _ ",
		" _| _| _||_||_ |_   ||_||_|",
		"  ||_  _|  | _||_|  ||_| _|",
	}
	//=> 123456789

	tc10 := []string{
		" _     _  _  _  _  _  _    ",
		"| || || || || || || ||_   |",
		"|_||_||_||_||_||_||_| _|  |",
	}
	//=> 000000051
	tc11 := []string{
		"    _  _  _  _  _  _     _ ",
		"|_||_|| ||_||_   |  |  | _ ",
		"  | _||_||_||_|  |  |  | _|",
	}
	//=> 490867715
	//fmt.Println(fixAccount(tc0))
	assert("711111111" == fixAccount(tc0), "test 16")
	assert("777777177" == fixAccount(tc1), "test 17")
	assert("200800000" == fixAccount(tc2), "test 18")
	assert("333393333" == fixAccount(tc3), "test 19")
	assert("888886888,888888988,888888880" == fixAccount(tc4), "test 20")
	assert("559555555,555655555" == fixAccount(tc5), "test 21")
	assert("686666666,666566666" == fixAccount(tc6), "test 22")
	assert("899999999,993999999,999959999" == fixAccount(tc7), "test 23")
	assert("490867715,490067115,490067719" == fixAccount(tc8), "test 24")
	assert("123456789" == fixAccount(tc9), "test 25")
	assert("000000051" == fixAccount(tc10), "test 26")
	assert("490867715" == fixAccount(tc11), "test 27")

}

func main() {
	TestStory1()
	TestStory3()
	TestStory4()
}
