package pointer

import "fmt"

func Pointer() {
	// var ptr *int  //just creating ptr
	// fmt.Println("value of pointer is", ptr)

	myNumber := 94
	var ptr = &myNumber //referencing also

	fmt.Println("value of actual ptr is", ptr)
	fmt.Println("value of actual ptr is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is ", myNumber) //188
	//fmt.Println("New value is ", *ptr)     //188

}
