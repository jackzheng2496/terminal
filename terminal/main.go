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
	reader   = bufio.NewReader(os.Stdin)
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
	for {
		if !loop {
			break
		}

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
				NewE := AddOnArgs(words[1], words)
				if NewE != nil {
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
		TODO: Start of reading in shittyDB
	*/
	terminalutil.LoadFromShittyDB(shittyDB, "./shittyDB.txt")
	/*
		End of reading in shittyDB
	*/
	RunningLoop()                                           //	Main loop of execution
	terminalutil.SaveToShittyDB(shittyDB, "./shittyDB.txt") //	Saving current info in shittyDB
}
