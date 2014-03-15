package main

import (
	"testing"
)

func assert(t bool) bool {
	if !t {
		return false
	}
	return true
}

// Test1() - standard test
func Test1(t *testing.T) {
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
	if actual, expect := b.String(), "Fo Fum"; !assert(actual == expect) {
		t.Fatalf("actual: %v, expect: %v", actual, expect)
	}
}

// Test2() - full buffer, no overlap
func Test2(t *testing.T) {
	b := NewBuffer(5)
	b.Add([]string{"one", "two", "three", "four", "five"})
	//b.List()
	if actual, expect := b.String(), "one two three four five"; !assert(actual == expect) {
		t.Fatalf("actual: %v, expect: %v", actual, expect)
	}
}

// Test3() - no overlap one element
func Test3(t *testing.T) {
	b := NewBuffer(5)
	b.Add([]string{"one", "two", "three", "four", "five", "six"})
	//b.List()
	if actual, expect := b.String(), "two three four five six"; !assert(actual == expect) {
		t.Fatalf("actual: %v, expect: %v", actual, expect)
	}
}

// Test4() - two full overlap
func Test4(t *testing.T) {
	b := NewBuffer(5)
	b.Add([]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"})
	if actual, expect := b.String(), "six seven eight nine ten"; !assert(actual == expect) {
		t.Fatalf("actual: %v, expect: %v", actual, expect)
	}
	b.Remove(5)
	if actual, expect := b.String(), ""; !assert(actual == expect) {
		t.Fatalf("actual: %v, expect: %v", actual, expect)
	}
}

// Test5() - empty
func Test5(t *testing.T) {
	b := NewBuffer(5)
	if actual, expect := b.String(), ""; !assert(actual == expect) {
		t.Fatalf("actual: %v, expect: %v", actual, expect)
	}
	b.Add([]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"})
	b.Add([]string{"eleven", "twelve", "thirteen", "fourteen", "fifteen"})
	b.Remove(4)
	if actual, expect := b.String(), "fifteen"; !assert(actual == expect) {
		t.Fatalf("actual: %v, expect: %v", actual, expect)
	}
}
