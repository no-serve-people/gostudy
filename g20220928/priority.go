package g20220928

import "fmt"

// Go语言在select语句中实现优先级
// https://www.liwenzhou.com/posts/Go/priority_in_go_select/
func worker2(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return

		case job1 := <-ch1:
			fmt.Println(job1)

		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority
				}
			}
			fmt.Println(job2)
		}
	}
}
