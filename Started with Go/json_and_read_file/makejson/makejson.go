package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil{
		fmt.Printf("Cannot read from STDIN %v", err.Error())
	}
	name = strings.TrimRight(name, "\r\n")
	fmt.Print("Enter your address: ")
	address, err := reader.ReadString('\n')
	if err != nil{
		fmt.Printf("Cannot read from STDIN %v", err.Error())
	}
	address = strings.TrimRight(address, "\r\n")
	person := Person{Name:name, Address: address}
	result, err := json.Marshal(person)
	if err != nil{
		fmt.Printf("Cannot marshall to JSON %v", err.Error())
	}
	fmt.Print(string(result))
}

type Person struct{
	Name string
	Address string
}
