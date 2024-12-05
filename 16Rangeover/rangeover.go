package rangeover

import "fmt"

func Rangeover() {

	nums := []int{2, 3, 5}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	for i, num := range nums { // [index , value]
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) //a->apple , b->banana
	}

	for k := range kvs { //only keys
		fmt.Println("key:", k)
	}

	for i, c := range "vo" {
		fmt.Println(i, c) //[0,118] [1,111]  asciis
	}

}
