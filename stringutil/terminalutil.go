package stringutil

import (
	"os"
)

func AddToMap(list map[string]Entertainer, e Entertainer, name string) {
	_, exist := list[name]
	if !exist {
		list[name] = e
	}
}

func GetMapValue(list map[string]Entertainer, name string) Entertainer {
	_, exist := list[name]
	if exist {
		return list[name]
	} else {
		return nil
	}
}

func RemoveMapValue(list map[string]Entertainer, name string) Entertainer {
	_, exist := list[name]
	if exist {
		temp := list[name]
		delete(list, name)
		return temp
	} else {
		return nil
	}
}

//TODO:	Figure out how to check type at runtime
func SaveType(e Entertainer, file *os.File) (n int, err error) {
	i, err := file.WriteString(e.LongOutput())
	return i, err
}
