package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input_text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Cannot read from Stdin %#v", err.Error())
	}
	res := strings.ToLower(input_text)
	res = strings.TrimSuffix(res, "\n")
	res = strings.TrimSuffix(res, "\r")
	if strings.ContainsAny(res,"ian") && strings.Index(res, "i") == 0 && strings.LastIndex(res, "n") == len(res)-1 {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
