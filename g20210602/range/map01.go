package main

import "fmt"

func main() {
	d := map[string]string{
		"asong": "帅",
		"song":  "太帅了",
	}

	for k := range d {
		if k == "asong" {
			delete(d, k)
		}
	}
	fmt.Println(d)
}
