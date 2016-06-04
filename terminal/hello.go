package main

import (
	"bufio"
	"fmt"
	"github.com/jackzheng2496/stringutil"
	"log"
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
	shittyDB = make(map[string]stringutil.Entertainer)
	reader   = bufio.NewReader(os.Stdin)
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
			return stringutil.NewManga(words[2], FormatTime, words[3], "0", "n/a")
		} else if len(words) == 3 {
			return stringutil.NewManga(words[2], FormatTime, "0", "0", "n/a")
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
					stringutil.AddToMap(shittyDB, NewE, words[2])
				} else {
					fmt.Println("Invalid Arguments")
				}
			}
		case LIST:
			if wordsLength == 2 {
				if strings.Compare(words[1], LONGLIST) == 0 {
					for _, value := range shittyDB {
						fmt.Println(value.LongOutput())
					}
				}
			} else {
				for _, value := range shittyDB {
					fmt.Println(value.FormattedOutput())
				}
			}
		case UPDATE:
			if wordsLength < 3 || wordsLength > 4 {
				fmt.Println("Too little arguments")
			} else {
				CurrentTime := time.Now()
				FormatTime := CurrentTime.Format(time.ANSIC)
				if wordsLength == 4 {
					if strings.Compare(words[1], SUBVAL) == 0 {
						stringutil.UpdateSub(shittyDB[words[2]], words[3])
						stringutil.UpdateTimestamp(shittyDB[words[2]], FormatTime)
					} else if strings.Compare(words[1], STUDIO) == 0 {
						stringutil.UpdatePublisher(shittyDB[words[2]], words[3])
						stringutil.UpdateTimestamp(shittyDB[words[2]], FormatTime)
					} else {
						fmt.Println("Invalid Arguments")
					}

				} else {
					stringutil.UpdateVal(shittyDB[words[1]], words[2])
					stringutil.UpdateTimestamp(shittyDB[words[1]], FormatTime)
				}

			}
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
}

func main() {
	/*
		TODO: Start of reading in shittyDB
	*/
	file, err := os.Open("./shittyDB.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 1000)
	data, num := stringutil.ReadPrevDB(file, data)

	buf := string(data[:num-1])
	index := strings.Split(buf, "\n")

	for val := range index {
		e, name := stringutil.CreateEntertainerFromLoad(index[val])
		stringutil.AddToMap(shittyDB, e, name)
	}

	/*
		End of reading in shittyDB
	*/

	RunningLoop() //	Main loop of execution
	fmt.Println("Saving to shittyDB...")
	SaveToShittyDB() //	Saving current info in shittyDB
}

func SaveToShittyDB() {
	file, err := os.Create("./shittyDB.txt")
	check(err)

	defer file.Close() //Idiomatic to defer file closing after opening

	for key, _ := range shittyDB {
		stringutil.SaveType(shittyDB[key], file)
		file.WriteString("\n")
	}

	file.Sync()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
