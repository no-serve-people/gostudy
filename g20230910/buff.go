package g20230910

import (
	"bufio"
	"fmt"
	"strings"
)

// https://www.v2ex.com/t/972407#reply3
func buff() {
	str := "12345678901234567890123456789012345678901234567890"

	strReader := strings.NewReader(str)
	bufReader := bufio.NewReaderSize(strReader, 16)
	p := make([]byte, 4)
	n, err := bufReader.Read(p)
	if err != nil {
		panic(err)
	}

	buffered := bufReader.Buffered()
	fmt.Printf("buffered:%d,content:%s\n", buffered, p[:n])

}
