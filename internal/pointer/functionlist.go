package pointer

func WithoutPointer(a int) {
	a = a + 100
}

func WithPointer(a *int) {
	*a = *a + 100
}
