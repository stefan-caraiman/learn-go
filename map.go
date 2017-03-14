package main

import "fmt"

func main() {
	presAge := make(map[string] int)
	presAge["Roosevlet"] = 42
	fmt.Println(presAge)
	presAge["Kenedy"] = 43
	fmt.Println(presAge["Kenedy"])
	delete(presAge, "Kenedy")
}

