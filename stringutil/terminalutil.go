package stringutil

func AddToMap(list map[string]Entertainment, e Entertainment, name string) {
	_, value := list[name]
	if !value {
		list[name] = e
	}
}

func GetMapValue(list map[string]Entertainment, name string) Entertainment {
	_, value := list[name]
	if value {
		return list[name]
	} else {
		return nil
	}
}

func RemoveMapValue(list map[string]Entertainment, name string) Entertainment {
	_, value := list[name]
	if value {
		temp := list[name]
		delete(list, name)
		return temp
	} else {
		return nil
	}
}
