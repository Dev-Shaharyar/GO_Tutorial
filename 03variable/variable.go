package variable

import "fmt"

func Variable() {

	//at the time of decl must specify the type
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e) //if not assign value then return 0

	f := "apple"
	fmt.Println(f)
}
