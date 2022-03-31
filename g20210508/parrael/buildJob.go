package main

import "fmt"

func buildJob(ch chan error, name string) {
	var err error

	fmt.Println(name)

	if _, ok := <-ch; !ok {
		return
	}
	ch <- err // finnaly, send the result into channel
}

func build() error {
	jobCount := 3
	errCh := make(chan error, jobCount)
	defer close(errCh) // 关闭 channel

	go buildJob(errCh, "A")
	go buildJob(errCh, "B")
	go buildJob(errCh, "C")

	for {
		select {
		case err := <-errCh:
			if err != nil {
				return err
			}
		}

		jobCount--
		if jobCount <= 0 {
			break
		}
	}

	return nil
}

func main() {
	build()
}
