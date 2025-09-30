package main

import (
    "fmt"
    "bufio"
    "os"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter a number: ")

	// input, err or _ if you can't handle error
	input, _ := reader.ReadString('\n')
	fmt.Println("you entered: ", input)
	fmt.Printf("you entered: %T", input)
}