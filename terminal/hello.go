package main

import (
	"bufio"
	"fmt"
	"github.com/jackzheng2496/stringutil"
	"os"
	"strings"
	"time"
)

const (
	PUT    string = "put"
	RM     string = "rm"
	UPDATE string = "update"
	GET    string = "get"
	LIST   string = "list"
	QUIT   string = "quit"
	ANIME  string = "-a"
	MANGA  string = "-m"
	SUBVAL string = "-s"
)

func AddOnArgs(option string, words []string) stringutil.Entertainer {
	CurrentTime := time.Now()
	FormatTime := CurrentTime.Format(time.ANSIC)

	switch option {
	case ANIME:
		if len(words) == 4 {
			return stringutil.NewAnime(words[2], FormatTime, words[3], "0", "n/a")
		} else if len(words) == 3 {
			return stringutil.NewAnime(words[2], FormatTime, "0", "0", "n/a")
		} else {
			return nil
		}
	case MANGA:
		if len(words) == 4 {
			return stringutil.NewManga(words[2], FormatTime, words[3], "0")
		} else if len(words) == 3 {
			return stringutil.NewManga(words[2], FormatTime, "0", "0")
		} else {
			return nil
		}
	default:
		return nil
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	shittyDB := make(map[string]stringutil.Entertainer)
	loop := true
	//	Main loop of terminal
	for {
		if !loop {
			break
		}

		fmt.Print("terminal: ")

		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")  //Remove newline character
		words := strings.Split(input, " ") //Split on space to get words

		switch words[0] { //Get first argument
		case PUT:
			if len(words) < 3 {
				fmt.Println("Too little arguments")
			} else {
				NewE := AddOnArgs(words[1], words)
				if NewE != nil {
					stringutil.AddToMap(shittyDB, NewE, words[2])
				} else {
					fmt.Println("Invalid Arguments")
				}
			}
		case LIST:
			for _, value := range shittyDB {
				value.FormattedOutput()
				fmt.Println()
			}
		case UPDATE:
			if len(words) < 3 || len(words) > 4 {
				fmt.Println("Too little arguments")
			} else {
				if len(words) == 4 {
					if strings.Compare(words[1], SUBVAL) == 0 {
						stringutil.UpdateSub(shittyDB[words[2]], words[3])
					} else {
						fmt.Println("Invalid Arguments")
					}
				} else {
					stringutil.UpdateVal(shittyDB[words[1]], words[2])
				}

			}
		case GET:
			if len(words) == 2 {
				_, exist := shittyDB[words[1]]
				if exist {
					shittyDB[words[1]].FormattedOutput()
					fmt.Println()
				} else {
					fmt.Println("No such key")
				}
			}
		case RM:
			if len(words) == 2 {
				_, exist := shittyDB[words[1]]
				if exist {
					stringutil.RemoveMapValue(shittyDB, words[1])
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
	fmt.Println("Saving to shittyDB...")
}
