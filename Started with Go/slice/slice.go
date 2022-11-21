package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main(){
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	result := make([]int, 0, 3)
	fmt.Print("Enter number: ")
	for scanner.Scan() {
		text = scanner.Text()
		if text != "X"{
			num, err := strconv.Atoi(text)
			if err != nil{
				fmt.Println("You used a letter, not a number -", err.Error())
				fmt.Print("Enter number: ")
			} else{
				result = append(result, num)
			sort.Slice(result, func(i, j int) bool {
				return result[i] < result[j]
			})
			fmt.Println("Result: ", result)
			fmt.Print("Enter number: ")
			}
		} else {
			break
		}
	}
	if err := scanner.Err(); err != nil{
		fmt.Printf("Cannot read from stdin %v", err.Error())
	}
}
