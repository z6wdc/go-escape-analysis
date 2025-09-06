package examples

func ReturnPointer() *int {
	x := 42
	return &x // returning a pointer to a local variable → escapes
}

func ReturnValue() int {
	x := 42
	return x // returning a local variable by value → does not escape
}
