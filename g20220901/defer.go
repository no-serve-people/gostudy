package g20220901

import (
	"log"
	"os"
)

func testDefer() {
	for i := 0; i < 5; i++ {
		func() {
			f, err := os.Open("/path/to/file")
			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()
		}()
	}
}
