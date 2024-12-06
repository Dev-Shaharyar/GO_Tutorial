package generics

import "fmt"

// func printslice(names []int) {
// 	for _, num := range nums {
// 		fmt.Println(num)
// 	}
// }

//	func printStringslice(names []string) {
//		for _, name := range names {
//			fmt.Println(name)
//		}
//	}
// func printslice[T any](nums []T) { //now it will works for any data type
// 	//we can also write the interface{} in place of any
// 	//and if u want to works for only int string the just put the int|string in place of the any
//  comparable can also be used

// 	for _, num := range nums {
// 		fmt.Println(num)
// 	}
// }

// type stack struct {
// 	elements []int
// }

type stack[T any] struct {
	elements []T
}

func Generics() {
	//nums := []int{1, 2, 3}
	// printslice(nums)

	//names := []string{"golang", "js", "java"}
	// printStringslice(names)

	// printslice(names)
	// printslice(nums)

	// myStack := stack{
	// 	elements: []int{1, 2, 3},
	// }

	myStack := stack[int]{
		elements: []int{1, 2, 3},
	}

	myStack2 := stack[string]{
		elements: []string{"go", "java"},
	}

	fmt.Println(myStack)

	fmt.Println(myStack2)

}
