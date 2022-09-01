package g20220901

import "fmt"

func testSelectFor2(chExit chan bool) {
EXIT:
	for {
		select {
		case val, ok := <-chExit:
			if !ok {
				fmt.Println("close channel 2", val)
				break EXIT
			}

			fmt.Println("ch2 val=", val)
		}
	}

	// EXIT2:
	fmt.Println("exit testSelectFor2")
}
