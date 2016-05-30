package stringutil

import (
	//"bytes"
	"fmt"
	//"strconv"
)

/*
	Some constants for available commands
*/
const (
	PUT    string = "put"
	RM     string = "rm"
	UPDATE string = "update"
	GET    string = "get"
	LIST   string = "list"
	QUIT   string = "quit"
)

func PrintList(e ...Entertainment) {
	for _, e := range e {
		e.FormattedOutput()
		fmt.Println()
	}
}

/*
	TODO: Figure out how to organize these structs more efficiently
*/

// func AddToList(list map[string]int, name string) int {
// 	_, value := list[name]
// 	if value == false {
// 		list[name] = 0
// 		return 0
// 	} else {
// 		return list[name]
// 	}
// }
//
// func UpdateListValue(list map[string]int, name string, value string) int {
// 	i, err := strconv.Atoi(value)
// 	if err != nil {
// 		fmt.Println(err)
// 		return -1
// 	}
//
// 	_, num := list[name]
// 	if num == false {
// 		return -1
// 	} else {
// 		list[name] = i
// 		return i
// 	}
// }
//
// func GetMapValue(list map[string]int, name string) int {
// 	_, num := list[name]
// 	if num != false {
// 		return list[name]
// 	} else {
// 		return 0
// 	}
// }
//
// func PrintList(list map[string]int) {
// 	fmt.Print(list)
// }
//
// func RemoveValue(list map[string]int, name string) {
// 	_, num := list[name]
// 	if num != false {
// 		delete(list, name)
// 	} else {
// 		fmt.Println("Invalid value")
// 	}
// }
