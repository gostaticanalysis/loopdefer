package a

func f1() {
	defer func() {}()
	for {
		defer func() {}() // want "defer should not use in a loop"
	}
}

func f2() {
loop:
	println()
	defer func() {}() // want "defer should not use in a loop"
	goto loop
}

func f3() {
loop:
	println()
	defer func() {}() // want "defer should not use in a loop"
	if false {
		goto loop
	}
}
