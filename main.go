package main

func main() {
	s := "world"
	for _, v := range []byte(s) {
		_ = v
	}
}
