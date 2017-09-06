package kiputil

func Foo() string {
	return "foo"
}

func Bar() string {
	return "bar"
}

func Baz() string {
	return "baz"
}

func Help() string {
	return "being helpful"
}

func Add(a int, b int) int {
	return a + b
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
