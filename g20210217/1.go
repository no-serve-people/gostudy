package main

import (
	"fmt"
	"net/url"
)

func main() {
	parse, _ := url.Parse("rtmps://live-api-s.facebook.com:443/rtmp/1318385771963181?s_bl=1&s_oil=0&s_psm=1&s_sw=0&s_tids=1&s_vt=api-s&a=AbyQIZiOeo12_otg")

	fmt.Println("scheme", parse.Scheme)
	fmt.Println("Opaque", parse.Opaque)
	fmt.Println("user", parse.User)
	fmt.Println("host", parse.Host)
	fmt.Println("Hostname", parse.Hostname())
	fmt.Println("Port", parse.Port())
	fmt.Println("path", parse.Path)
	fmt.Println("rawpath", parse.RawPath)
	fmt.Println("forcequery", parse.ForceQuery)
	fmt.Println("rawquery", parse.RawQuery)
	fmt.Println("Fragment", parse.Fragment)
	fmt.Println("RawFragment", parse.RawFragment)
}
