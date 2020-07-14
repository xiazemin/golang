package main

import "github.com/armon/go-radix"

func main(){
	// Create a tree
	r := radix.New()
	r.Insert("foo", 1)
	r.Insert("bar", 2)
	r.Insert("foobar", 2)

	// Find the longest prefix match
	m, _, _ := r.LongestPrefix("foozip")
	if m != "foo" {
		panic("should be foo")
	}
}