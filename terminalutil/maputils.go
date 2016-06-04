package terminalutil

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
