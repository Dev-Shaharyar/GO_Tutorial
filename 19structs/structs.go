package structs

import "fmt"

func Structs() {
	//no inharitance in golang ; no super or parent
	//fmt.Println("hello")
	alam := User{"shaharyar", "shaharyar@insta", true, 22}
	fmt.Println(alam)
	fmt.Printf("shaharyar details are:%+v\n", alam)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
