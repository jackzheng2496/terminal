package main

import (
	"bufio"
	"fmt"
	"github.com/jackzheng2496/terminalutil"
	"os"
	"strings"
	"time"
)

const (
	//	TODO:	Consider using "flags" to parse instead of hand parsing it
	PUT      string = "put"
	RM       string = "rm"
	UPDATE   string = "update"
	GET      string = "get"
	LIST     string = "list"
	LONGLIST string = "-l"
	QUIT     string = "quit"
	ANIME    string = "-a"
	MANGA    string = "-m"
	SUBVAL   string = "-s"
	STUDIO   string = "-st"
)

var (
	shittyDB = make(map[string]terminalutil.Entertainer)
	Auth     = make(map[string]terminalutil.User)
	//	TODO:	Consider using NewScanner instead
	reader = bufio.NewReader(os.Stdin)
)

func AddOnArgs(option string, words []string) terminalutil.Entertainer {
	CurrentTime := time.Now()
	FormatTime := CurrentTime.Format(time.ANSIC)

	switch option {
	case ANIME:
		if len(words) == 4 {
			return terminalutil.NewAnime(words[2], FormatTime, words[3], "0", "n/a")
		} else if len(words) == 3 {
			return terminalutil.NewAnime(words[2], FormatTime, "0", "0", "n/a")
		} else {
			return nil
		}
	case MANGA:
		if len(words) == 4 {
			return terminalutil.NewManga(words[2], FormatTime, words[3], "0", "n/a")
		} else if len(words) == 3 {
			return terminalutil.NewManga(words[2], FormatTime, "0", "0", "n/a")
		} else {
			return nil
		}
	default:
		return nil
	}

}

func RunningLoop() {
	loop := true
	for loop {
		fmt.Print("terminal: ")

		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")  //Remove newline character
		words := strings.Split(input, " ") //Split on space to get words

		wordsLength := len(words)

		switch words[0] { //Get first argument
		case PUT:
			if wordsLength < 3 {
				fmt.Println("Too little arguments")
			} else {
				if NewE := AddOnArgs(words[1], words); NewE != nil {
					terminalutil.AddToMap(shittyDB, NewE, words[2])
				} else {
					fmt.Println("Invalid Arguments")
				}
			}
		case LIST:
			if wordsLength == 2 {
				if strings.Compare(words[1], LONGLIST) == 0 {
					terminalutil.PrettyPrintingLong(shittyDB)
				}
			} else {
				terminalutil.PrettyPrintingShort(shittyDB)
			}
		case UPDATE:
			if wordsLength < 3 || wordsLength > 4 {
				fmt.Println("Too little arguments")
			} else {
				CurrentTime := time.Now()
				FormatTime := CurrentTime.Format(time.ANSIC)
				if wordsLength == 4 {
					if strings.Compare(words[1], SUBVAL) == 0 {
						terminalutil.UpdateSub(shittyDB[words[2]], words[3])
						terminalutil.UpdateTimestamp(shittyDB[words[2]], FormatTime)
					} else if strings.Compare(words[1], STUDIO) == 0 {
						terminalutil.UpdatePublisher(shittyDB[words[2]], words[3])
						terminalutil.UpdateTimestamp(shittyDB[words[2]], FormatTime)
					} else {
						fmt.Println("Invalid Arguments")
					}

				} else {
					terminalutil.UpdateVal(shittyDB[words[1]], words[2])
					terminalutil.UpdateTimestamp(shittyDB[words[1]], FormatTime)
				}

			}
		/*
			TODO:	Is there a point of GET if list does almost the same thing?
		*/
		case GET:
			if wordsLength == 2 {
				_, exist := shittyDB[words[1]]
				if exist {
					fmt.Println(shittyDB[words[1]].FormattedOutput())
				} else {
					fmt.Println("No such key")
				}
			}
		case RM:
			if wordsLength == 2 {
				_, exist := shittyDB[words[1]]
				if exist {
					terminalutil.RemoveMapValue(shittyDB, words[1])
				} else {
					fmt.Println("No such key")
				}
			}
		case QUIT:
			loop = false
		default:
			fmt.Println("Invalid Command")
		}
	}
}

func main() {
	/*
		TODO: User Authentication -> see if a file exists for them
	*/
	fmt.Printf("Username: ")
	input, _ := reader.ReadString('\n')
	username := strings.Trim(input, "\n") //Remove newline character
	fmt.Printf("Password: ")
	input, _ = reader.ReadString('\n')
	password := strings.Trim(input, "\n")

	up := &terminalutil.User{Username: username, Password: password, Filename: "./" + username + ".txt"}
	file, err := os.Open("./userpass.txt")
	if err != nil {
		fmt.Println("Does not exist")
	}

	terminalutil.FillMapWithID(file, Auth)
	bool := terminalutil.CheckAuthentication(up, Auth)
	file.Close()
	if bool {
		/*
			TODO: Start of reading in shittyDB
		*/
		terminalutil.LoadFromShittyDB(shittyDB, up.Filename)
	} else {
		fmt.Printf("Create account? (Y/N) ")
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		switch input {
		case "Y":
			Auth[up.Username] = *up
		case "N":
			fmt.Println("Exiting app...")
			os.Exit(0)
		}
	}
	/*
		End of reading in shittyDB
	*/
	RunningLoop() //	Main loop of execution
	//	TODO:	Consider saving in JSON format
	terminalutil.SaveToShittyDB(shittyDB, up.Filename) //	Saving current info in shittyDB
	/*
		TODO: Save User-Pass combination in a file
	*/
	file, err = os.OpenFile("./userpass.txt", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	terminalutil.SaveMapWithID(file, Auth)
}
