package terminalutil

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadPrevDB(file *os.File, buffer []byte) ([]byte, int) {
	i, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	return buffer, i
}

/*
	TODO: 	Save data as JSON format after getting User Authentication to work
*/
func SaveToShittyDB(shittyDB map[string]Entertainer, filename string) {
	fmt.Println("Saving to", filename, "...")
	file, err := os.Create(filename)
	check(err)
	defer file.Close() //Idiomatic to defer file closing after opening

	for key, _ := range shittyDB {
		SaveType(shittyDB[key], file)
		file.WriteString("\n")
	}
	file.Sync()
}

func LoadFromShittyDB(shittyDB map[string]Entertainer, filename string) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	data := make([]byte, 1000)
	data, num := ReadPrevDB(file, data)

	buf := string(data[:num-1])
	index := strings.Split(buf, "\n")

	for val := range index {
		e, name := CreateEntertainerFromLoad(index[val])
		AddToMap(shittyDB, e, name)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
