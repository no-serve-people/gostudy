package g20220901

func testRecover1() {
	recover()
	panic(1)
}

func testRecover2() {
	defer recover()
	panic(1)
}

func testRecover3() {
	defer func() {
		func() {
			recover()
		}()
	}()
	panic(1)
}

func testRecover4() {
	defer func() {
		recover()
	}()

	panic(1)
}
