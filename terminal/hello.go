package main

import (
	"bufio"
	"fmt"
	"github.com/jackzheng2496/stringutil"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data := make(map[string]int)

	for {
		fmt.Printf("$>jsh: ")

		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		splitArgs := strings.Split(text, " ")

		switch splitArgs[0] {
		case stringutil.PUT:
			if len(splitArgs) == 3 {
				stringutil.AddToList(data, splitArgs[1])
				stringutil.UpdateListValue(data, splitArgs[1], splitArgs[2])
			} else if len(splitArgs) == 2 {
				stringutil.AddToList(data, splitArgs[1])
			} else {
				fmt.Println("Missing Arguments")
			}
		case stringutil.UPDATE:
			if len(splitArgs) == 3 {
				stringutil.UpdateListValue(data, splitArgs[1], splitArgs[2])
			} else {
				fmt.Println("Missing Arguments")
			}
		case stringutil.LIST:
			stringutil.PrintList(data)
			fmt.Print("\n")
		case stringutil.GET:
			if len(splitArgs) == 2 {
				val := stringutil.GetMapValue(data, splitArgs[1])
				fmt.Println(splitArgs[1], ":", val)
			} else {
				fmt.Println("Missing Arguments")
			}
		case stringutil.RM:
			if len(splitArgs) == 2 {
				stringutil.RemoveValue(data, splitArgs[1])
			} else {
				fmt.Println("Missing Arguments")
			}
		case stringutil.QUIT:
			fmt.Println("Exiting app...")
			os.Exit(0)
		default:
			fmt.Println("No such Command")
		}

	}
}
