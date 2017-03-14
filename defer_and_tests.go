package main

import "fmt"

func main() {
	// defer executes at the end of the run
	defer printTwo()
	printOne()
	defer printTwo()
	printOne()

	fmt.Println(safeDiv(3, 0))
}

func safeDiv(num1, num2 int) int{
	defer func(){
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution	
}

func printOne() { fmt.Println(1)}
func printTwo() { fmt.Println(2)}