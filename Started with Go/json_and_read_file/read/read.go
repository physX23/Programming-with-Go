package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parse_file(file string) []string {
	res := make([]string, 0)
	data, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error when open %v", err.Error())
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fname := strings.Split(scanner.Text(), " ")[0]
		lname := strings.Split(scanner.Text(), " ")[1]
		if len(fname) > 20 {
			fname = fname[:20]
		}
		if len(lname) > 20 {
			lname = lname[:20]
		}
		result := fname + " " + lname
		res = append(res, result)
	}
	return res
}

func main() {
	fmt.Print("File path:\n")
	reader := bufio.NewReader(os.Stdin)
	file, err := reader.ReadString('\n')
	file = strings.TrimRight(file, "\r\n")
	if err != nil{
		fmt.Printf("Cannot read from STDIN %v", err.Error())
	}
	result := parse_file(file)
	for _, value := range result {
		fmt.Println(value)
	}

}
