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
)

func AddOnArgs(option string, words []string) stringutil.Entertainment {
	CurrentTime := time.Now()
	FormatTime := CurrentTime.Format(time.ANSIC)

	switch option {
	case ANIME:
		if len(words) == 4 {
			return stringutil.NewAnime(words[2], FormatTime, words[3], "0")
		} else if len(words) == 3 {
			return stringutil.NewAnime(words[2], FormatTime, "0", "0")
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
	shittyDB := make(map[string]stringutil.Entertainment)
	//	Main loop of terminal
	for {
		fmt.Print("terminal: ")

		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")  //Remove newline character
		words := strings.Split(input, " ") //Split on space to get words

		switch words[0] { //Get first argument
		case PUT:
			NewE := AddOnArgs(words[1], words)
			if NewE != nil {
				stringutil.AddToMap(shittyDB, NewE, words[2])
			} else {
				fmt.Println("Invalid Arguments")
			}
		case LIST:
			for _, value := range shittyDB {
				stringutil.PrintList(value)
			}
		}

	}
	// s := time.Now()
	// formatTime := s.Format(time.ANSIC)
	// testAnime := stringutil.NewAnime("name", formatTime, 5, 2)
	// formatTime = s.Format(time.ANSIC)
	// testManga := stringutil.NewManga("Manga", formatTime, 4, 1)
	// stringutil.PrintList(testAnime, testManga)
	// fmt.Println()

	//reader := bufio.NewReader(os.Stdin)
	//data := make(map[string]int)

	// for {
	// 	fmt.Printf("$>jsh: ")
	//
	// 	text, _ := reader.ReadString('\n')
	// 	text = text[:len(text)-1]
	// 	splitArgs := strings.Split(text, " ")
	//
	// 	switch splitArgs[0] {
	// 	case stringutil.PUT:
	// 		if len(splitArgs) == 3 {
	// 			stringutil.AddToList(data, splitArgs[1])
	// 			stringutil.UpdateListValue(data, splitArgs[1], splitArgs[2])
	// 		} else if len(splitArgs) == 2 {
	// 			stringutil.AddToList(data, splitArgs[1])
	// 		} else {
	// 			fmt.Println("Missing Arguments")
	// 		}
	// 	case stringutil.UPDATE:
	// 		if len(splitArgs) == 3 {
	// 			stringutil.UpdateListValue(data, splitArgs[1], splitArgs[2])
	// 		} else {
	// 			fmt.Println("Missing Arguments")
	// 		}
	// 	case stringutil.LIST:
	// 		stringutil.PrintList(data)
	// 		fmt.Print("\n")
	// 	case stringutil.GET:
	// 		if len(splitArgs) == 2 {
	// 			val := stringutil.GetMapValue(data, splitArgs[1])
	// 			fmt.Println(splitArgs[1], ":", val)
	// 		} else {
	// 			fmt.Println("Missing Arguments")
	// 		}
	// 	case stringutil.RM:
	// 		if len(splitArgs) == 2 {
	// 			stringutil.RemoveValue(data, splitArgs[1])
	// 		} else {
	// 			fmt.Println("Missing Arguments")
	// 		}
	// 	case stringutil.QUIT:
	// 		fmt.Println("Exiting app...")
	// 		os.Exit(0)
	// 	default:
	// 		fmt.Println("No such Command")
	// 	}
	//
	// }
}
