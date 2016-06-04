package stringutil

import (
	"log"
	"os"
	"strings"
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

/*
	TODO:	**** Figure out how to check type at runtime ****
*/
func SaveType(e Entertainer, file *os.File) (n int, err error) {
	switch e.GetType() {
	case "Anime":
		file.WriteString("a ")
	case "Manga":
		file.WriteString("m ")
	}
	num, err := file.WriteString(e.SaveOutput())
	return num, err
}

func ReadPrevDB(file *os.File, buffer []byte) ([]byte, int) {
	i, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	return buffer, i
}

func CreateEntertainerFromLoad(name string) (Entertainer, string) {
	split := strings.Split(name, " ")
	if strings.Compare(split[0], "a") == 0 {
		return NewAnime(split[1], strings.Join(split[5:], " "), split[3], split[2], split[4]), split[1]
	} else {
		return NewManga(split[1], strings.Join(split[5:], " "), split[3], split[2], split[4]), split[1]
	}
}
