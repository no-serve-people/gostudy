package main

import (
	"testing"
)

type Item struct {
	Id   int64
	Name string
}

func BenchmarkRangeSpeed01(b *testing.B) {
	var items [1024]Item

	for i := 0; i < b.N; i++ {
		var tmp int64
		for _, item := range items {
			tmp = item.Id
		}
		// fmt.Println(tmp)
		_ = tmp
	}
}

func BenchmarkRangeSpeed02(b *testing.B) {
	var items [1024]Item

	for i := 0; i < b.N; i++ {
		var tmp int64
		for k := range items {
			item := items[k]
			tmp = item.Id
		}
		// fmt.Println(tmp)
		_ = tmp
	}
}
