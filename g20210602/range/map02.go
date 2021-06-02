package main

import "fmt"

func main() {
	d := map[string]string{
		"asong": "帅",
		"song":  "太帅了",
	}

	for k, v := range d {
		d[v] = k
	}
	fmt.Println(d)
}
