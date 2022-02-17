package main

func main() {
	a := make(chan bool, 100)
	b := make(chan bool, 100)
	c := make(chan bool, 100)

	for i := 0; i < 10; i++ {
		a <- true
		b <- true
		c <- true
	}

	for i := 0; i < 10; i++ {
		select {
		case <-a:
			print("< a")
		case <-b:
			print("< b")
		case <-c:
			print("< c")

		default:
			print("< default")
		}
	}
}
