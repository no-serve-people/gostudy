package _struct

import "fmt"

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func testSet() {
	s := make(Set)
	s.Add("TOM")
	s.Add("SAM")
	fmt.Println(s.Has("TOM"))
	fmt.Println(s.Has("JA"))
}
