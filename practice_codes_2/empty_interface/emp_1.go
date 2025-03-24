package main

func main_1() {
	var i int8 = 1
	read(i)
}

//go:noinline
func read(i interface{}) {

	n := i.(int16)
	println(n)
}
