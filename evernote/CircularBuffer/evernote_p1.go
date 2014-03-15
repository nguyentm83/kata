// https://evernote.com/careers/challenge.php
// Circular buffer
package main

import (
	"flag"
	"fmt"
	"strings"
)

// Circular buffer
type Buffer struct {
	buffer  []string // buffer
	count   uint     // element count
	in, out uint     // in, out pointer
}

func NewBuffer(size uint) *Buffer {
	b := Buffer{buffer: make([]string, size), count: 0, in: 0, out: 0}
	return &b
}

func (b *Buffer) Add(ele []string) {
	for _, v := range ele {
		b.buffer[b.in] = v
		b.in++
		if b.in >= b.Size() {
			b.in -= b.Size()
		}
	}

	if newcount := b.count + uint(len(ele)); newcount <= b.Size() {
		b.count = newcount
	} else {
		b.count = b.Size()
		b.out = b.in
	}
}

func (b *Buffer) Remove(count uint) {
	// invariant: count <= b.count
	b.count -= count
	b.out += count
	for b.out >= b.Size() {
		b.out -= b.Size()
	}
}

func (b *Buffer) Dump() {
	fmt.Printf("%v\n", b)
}

// implement Stringer interface for unit testing
func (b *Buffer) String() string {
	s := make([]string, b.count)
	for i, c := b.out, uint(0); c < b.count; i, c = i+1, c+1 {
		if i >= b.Size() {
			i -= b.Size()
		}
		s[c] = b.buffer[i]
	}
	return strings.Join(s, " ")
}

func (b *Buffer) List() {
	for i, c := b.out, uint(0); c < b.count; i, c = i+1, c+1 {
		if i >= b.Size() {
			i -= b.Size()
		}
		fmt.Printf("%v\n", b.buffer[i])
	}
}

func (b *Buffer) Size() uint {
	return uint(len(b.buffer))
}

func main() {
	var testmode = flag.Bool("test", false, "test mode")
	flag.Parse()
	if *testmode {
		fmt.Println("*** TEST MODE ***")
		test1()
		test2()
		test3()
		test4()
		test5()
		fmt.Println("*** END TEST ***")
	} else {
		run()
	}
}

func run() {
	var bufSize uint
	fmt.Scanf("%d\n", &bufSize)
	//fmt.Printf("BufSize = %v\n", bufSize)
	b := NewBuffer(bufSize)

	var cmd, op string
	for {
		fmt.Scanf("%s", &cmd)
		if cmd == "A" {
			fmt.Scanf("%s", &op)

			var n int
			var s string
			var list []string
			fmt.Sscanf(op, "%d", &n)
			list = make([]string, n)
			for i := 0; i < n; i++ {
				fmt.Scanf("%s", &s)
				list[i] = s
			}

			//fmt.Printf("Add(%v)\n", list)
			b.Add(list)
			continue
		}

		if cmd == "R" {
			fmt.Scanf("%s", &op)
			var n uint
			fmt.Sscanf(op, "%d", &n)
			//fmt.Printf("Remove(%v)\n", n)
			b.Remove(n)
			continue
		}

		if cmd == "L" {
			b.List()
		}

		if cmd == "Q" {
			break
		}
	}
}

func test(test, actual, expect string) bool {
	if actual != expect {
		fmt.Printf("%s FAILED!\n", test)
		fmt.Printf("actual: %v, expect: %v\n", actual, expect)
		return false
	}
	fmt.Printf("%s PASSED!\n", test)
	return true
}

//func assert(t bool, note string) bool {
//	if !t {
//		log.Printf("%s FAILED !\n", note)
//		return false
//	}
//	log.Printf("%s PASS !\n", note)
//	return true
//}

// standard test
func test1() {
	/*
		10
		A 3
		Fee
		Fi
		Fo
		A 1
		Fum
		R 2
		L
		Q*/
	b := NewBuffer(10)
	b.Add([]string{"Fee", "Fi", "Fo"})
	b.Add([]string{"Fum"})
	b.Remove(2)
	//b.List()
	test("test1", b.String(), "Fo Fum")
}

// test2() - full buffer, no overlap
func test2() {
	b := NewBuffer(5)
	b.Add([]string{"one", "two", "three", "four", "five"})
	//b.List()
	test("test2", b.String(), "one two three four five")
}

// test3() - no overlap one element
func test3() {
	b := NewBuffer(5)
	b.Add([]string{"one", "two", "three", "four", "five", "six"})
	//b.List()
	test("test3", b.String(), "two three four five six")
}

// test4() - two full overlap
func test4() {
	b := NewBuffer(5)
	b.Add([]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"})
	test("test4a", b.String(), "six seven eight nine ten")
	b.Remove(5)
	test("test4b", b.String(), "")
}

// test5() - empty
func test5() {
	b := NewBuffer(5)
	test("test5a", b.String(), "")
	b.Add([]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"})
	b.Add([]string{"eleven", "twelve", "thirteen", "fourteen", "fifteen"})
	b.Remove(4)
	test("test5b", b.String(), "fifteen")
}
