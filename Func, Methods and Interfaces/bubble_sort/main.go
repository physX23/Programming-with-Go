package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var notSorted bool

func BubbleSort (sli []int) {
	notSorted = true
	for notSorted {
		notSorted = false
		for i := range sli {
			if i < len(sli)-1 {
				Swap(sli, i)
			}
		}
	}
	fmt.Println("Result:")
	fmt.Println(sli)
}

func Swap (sli []int, idx int) {
	if sli[idx] > sli[idx+1] {
		sli[idx], sli[idx+1] = sli[idx+1], sli[idx]
		notSorted = true
	}
}

func main () {
	slice := make([]int, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter up to 10 numbers.")
	fmt.Println("Type 'q' if you want to stop typing\n")
	for len(slice) < 10 {
		fmt.Print("Enter number: ")
		text, err := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if err != nil {
			fmt.Printf("Cannot read from stdin %v", err.Error())
		} else if text == "q" {
			break
		} else {
			num, err := strconv.Atoi(text)
			if err != nil{
				fmt.Println("You used a letter, not a number -", err.Error())
				fmt.Print("Enter number: ")
			}
			slice = append(slice, num)
		}
	}
	BubbleSort(slice)
}
