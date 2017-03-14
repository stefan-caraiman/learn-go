package main

import "fmt"


func main() {
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	for j := 0; j < 5; j++ {
		fmt.Println(j) //0..4
	}
	age := 18
	if age >= 16 {
		fmt.Println("bye")
	} else if age == 18 {
		fmt.Println("hello")
	} else {
		fmt.Println("nope")
	}
	var favNums[5] float64
	favNums[0] = 3.14
	fmt.Println(favNums[0])
	favNumsOther := [5] int {1,2,3,4,5}
	for i, value := range favNumsOther {
		fmt.Println(value, i)
	}
	//slicing
	numSlice := []int {6, 7, 8, 9}
	numSlice2 := numSlice[2:4]
	fmt.Println(numSlice2)
	numSlice3 := make([]int, 5, 10)
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3)
}