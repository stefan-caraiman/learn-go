package main

import "fmt"

func addUp(numbers []float64) float64{
	var sum float64 = 0 // sum := 0.0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func main() {
	listNums := []float64{1,2,3,4,5}
	fmt.Println(addUp(listNums))
	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)
	fmt.Println(substractThem(1,2,3,4))

	//functions within functions
	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		return num3
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
	//calling recursive function
	fmt.Println(factorial(5))
}

func next2Values(number int) (int, int) {
	return number+1, number+2
}

func substractThem(int_args ...int) int {
	finalValue := 0
	for _, value := range int_args {
			finalValue -= value
	}
	return finalValue
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num - 1)
}
