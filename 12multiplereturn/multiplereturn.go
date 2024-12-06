package multiplereturn

import "fmt"

func vals() (int, int, int) {
	return 3, 7, 10
}

func Multiplereturn() {

	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	_, _, e := vals()
	//fmt.Println(d)
	fmt.Println(e)
}
