package limit

//
//func Test_Counter(t *testing.T) {
//	counter := NewCounter(0, 5, 100, time.Now().Unix())
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			for k := 0; k <= 10; k++ {
//				fmt.Println(counter.RateLimit())
//				if k%3 == 0 {
//					time.Sleep(102 * time.Millisecond)
//				}
//			}
//		}(i)
//		time.Sleep(10 * time.Second)
//	}
//}
