package terminalutil

import (
	"log"
	"os"
	"strings"
)

type User struct {
	Username, Password, Filename string
}

/*
   TODO:   Read from textfile containing user-pass combinations
           and check to see if given user-pass exists

           Yes:    Read in user specific file
           No:     Prompt user to make an account
*/
func CheckAuthentication(user *User, list map[string]User) bool {
	_, exist := list[user.Username]
	if exist {
		if strings.Compare(list[user.Username].Password, user.Password) == 0 {
			return true
		}
	}
	return false
}

func FillMapWithID(file *os.File, list map[string]User) {
	buffer := make([]byte, 1000)
	_, err := file.Read(buffer)

	if err != nil {
		log.Fatal(err)
	}

	buf := string(buffer[:len(buffer)-1])
	userpass := strings.Split(buf, "\n")

	//      Read in file as User structs
	for value := range userpass {
		up := strings.Split(userpass[value], " ")
		if len(up) > 1 {
			list[up[0]] = User{Username: up[0], Password: up[1]}
		}
	}
}

func SaveMapWithID(file *os.File, list map[string]User) {

}
