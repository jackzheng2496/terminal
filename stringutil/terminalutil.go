package stringutil

import (
	"fmt"
	"strconv"
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

/*
	Decided to make a struct to contain values instead of
	key:value map
*/
type Entertainment struct {
	name, kind string
	timestamp  float64
	count      int
}

/*
	Not too sure how this works, so UpdateCount can only
	be called with Entertainment structs?
*/
func (e *Entertainment) UpdateCount(count string) int {
	i, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	e.count = i
	return i
}

func (e *Entertainment) GetCount() int {
	return e.count
}

func AddToList(list map[string]Entertainment, e Entertainment) int {
	_, value := list[e.name]
	if value == false {
		list[e.name] = e
		return e.count
	} else {
		return list[e.name].count
	}
}

func RemoveKey(list map[string]Entertainment, name string) {
	_, e := list[name]
	if e != false {
		delete(list, name)
	} else {
		fmt.Println("Invalid value")
	}
}

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
