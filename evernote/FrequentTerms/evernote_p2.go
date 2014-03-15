// https://evernote.com/careers/challenge.php
// Frequent Terms

package main

import (
	"fmt"
	"sort"
)

type Count struct {
	w string // word
	c int    //count
}

type ByFreq []*Count

func (f ByFreq) Len() int           { return len(f) }
func (f ByFreq) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ByFreq) Less(i, j int) bool { return f[i].c < f[j].c || (f[i].c == f[j].c && f[i].w > f[j].w) }

var wordcount, k int
var book map[string]*Count
var list []*Count

func init() {
	book = make(map[string]*Count)
}

func run() {
	var word string
	fmt.Scanf("%d\n", &wordcount)
	for i := 0; i < wordcount; i++ {
		fmt.Scanf("%s\n", &word)
		process(word)
	}
	fmt.Scanf("%d\n", &k)
	sort.Sort(ByFreq(list))
	for i, j := 0, len(list)-1; i < k && j >= 0; i, j = i+1, j-1 {
		fmt.Printf("%v\n", list[j].w)
	}
}

func process(w string) {
	if ele, ok := book[w]; ok {
		ele.c += 1
	} else {
		book[w] = &Count{w: w, c: 1}
		list = append(list, book[w])
	}
}

func main() {
	run()
}
