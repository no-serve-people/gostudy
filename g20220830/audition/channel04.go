package main

// @控制并发数

// https://topgoer.cn/docs/goquestions/goquestions-1cjh2jsk14499

//var limit = make(chan int, 3)
//
//func main() {
//	for _, w := range worker {
//		go func() {
//			limit <- 1
//			w()
//			<-limit
//		}()
//	}
//}
