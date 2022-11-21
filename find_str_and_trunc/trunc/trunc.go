package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var index int
	var full_str string
	var result int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input_text, err := reader.ReadString('\n')
	if err != nil{
		fmt.Printf("Cannot read from Stdin %#v", err.Error())
	}
	if strings.Contains(input_text, "."){
		index = strings.Index(input_text, ".")
		full_str = input_text[:index]
		result, _ = strconv.Atoi(full_str)
	} else if strings.Contains(input_text, ","){
		index = strings.Index(input_text, ",")
		full_str = input_text[:index]
		result, _ = strconv.Atoi(full_str)
	}
	if result != 0 {
		fmt.Printf("Result: %v", result)
	} else if result == 0 && index == 1{
		fmt.Printf("Result: %v", result)
	} else{
		fmt.Printf("Entered value is not float")
	}
}