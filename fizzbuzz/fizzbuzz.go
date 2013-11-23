// http://codingdojo.org/cgi-bin/wiki.pl?back=KataFizzBuzz
package main

import (
	"fmt"
	"log"
	"strings"
)

type Test struct {
	q int    // input
	a string // expected output
}

func assert(t bool, note string) {
	if !t {
		log.Fatalf("%s FAILED !", note)
	}
}

func Game1(i int) string {
	// if a number is divisible by 3 => fizz
	// if a number is divisible by 5 => buzz
	// if a number is divisible by 3 and 5 => fizzbuzz
	if i%15 == 0 {
		return "FizzBuzz"
	}

	if i%3 == 0 {
		return "Fizz"
	}

	if i%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprintf("%d", i)
}

func Game2(i int) string {
	// A number is fizz if it is divisible by 3 or if it has a 3 in it
	// A number is buzz if it is divisible by 5 or if it has a 5 in it

	t := fmt.Sprintf("%d", i)
	if i%15 == 0 || (strings.Contains(t, "3") && strings.Contains(t, "5")) {
		return "FizzBuzz"
	}

	if i%3 == 0 || strings.Contains(t, "3") {
		return "Fizz"
	}

	if i%5 == 0 || strings.Contains(t, "5") {
		return "Buzz"
	}

	return fmt.Sprintf("%d", i)
}

func TestGame1() {
	log.Printf("GAME1 TEST STARTED ...")
	test := []Test{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
	}

	for _, v := range test {
		assert(v.a == Game1(v.q), "")
	}
}

func TestGame2() {
	log.Printf("GAME 2 TEST STARTED ...")
	test := []Test{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{31, "Fizz"},
		{53, "FizzBuzz"},
		{75, "FizzBuzz"},
		{30, "FizzBuzz"},
	}

	for i, v := range test {
		note := fmt.Sprintf("Test %d", i)
		assert(v.a == Game2(v.q), note)
	}
	Game2(53)

}

func main() {
	TestGame1()
	TestGame2()
	//fmt.Println("helllo")
}
