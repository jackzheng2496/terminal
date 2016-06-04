package terminalutil

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type User struct {
	Username, Password, filename string
}

/*
   TODO:   Read from textfile containing user-pass combinations
           and check to see if given user-pass exists

           Yes:    Read in user specific file
           No:     Prompt user to make an account
*/
func CheckAuthentication(user *User, file *os.File) bool {
	buffer := make([]byte, 1000)
	_, err := file.Read(buffer)

	if err != nil {
		log.Fatal(err)
	}

	list := make(map[string]User)
	buf := string(buffer[:len(buffer)-1])
	userpass := strings.Split(buf, "\n")

	//      Read in file as User structs
	for value := range userpass {
		ele := userpass[value]
		up := strings.Split(ele, " ")
		if len(up) > 1 {
			fmt.Println(up)
			u := User{Username: up[0], Password: up[1]}
			list[up[0]] = u
		}

	}

	_, exist := list[user.Username]

	if exist {
		if strings.Compare(list[user.Username].Password, user.Password) == 0 {
			return true
		}
	}
	return false
}
