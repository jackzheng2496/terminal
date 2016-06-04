package terminalutil

import (
	"fmt"
	"os"
	"strings"
)

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

func CreateEntertainerFromLoad(name string) (Entertainer, string) {
	split := strings.Split(name, " ")
	if strings.Compare(split[0], "a") == 0 {
		return NewAnime(split[1], strings.Join(split[5:], " "), split[3], split[2], split[4]), split[1]
	} else {
		return NewManga(split[1], strings.Join(split[5:], " "), split[3], split[2], split[4]), split[1]
	}
}

func PrettyPrintingShort(list map[string]Entertainer) {
	/*
		Short listing for animes
	*/
	fmt.Printf("\n%-20s%-10s%-10s\n", "Name", "Season", "Episode")
	fmt.Printf("-------------------------------------\n")
	for _, value := range list {
		if strings.Compare(value.GetType(), "Anime") == 0 {
			val := strings.Split(value.FormattedOutput(), " ")
			fmt.Printf("%-20s%-10s%-10s\n", val[0], val[1], val[2])
		}
	}
	fmt.Println()

	/*
		Short listing for mangas
	*/
	fmt.Printf("%-20s%-10s%-10s\n", "Name", "Volume", "Chapter")
	fmt.Printf("-------------------------------------\n")
	for _, value := range list {
		if strings.Compare(value.GetType(), "Manga") == 0 {
			val := strings.Split(value.FormattedOutput(), " ")
			fmt.Printf("%-20s%-10s%-10s\n", val[0], val[1], val[2])
		}
	}
	fmt.Println()
}

func PrettyPrintingLong(list map[string]Entertainer) {
	/*
		Long listing for animes
	*/
	fmt.Printf("\n%-20s%-10s%-10s%-15s%-15s\n", "Name", "Season", "Episode", "Studio", "Modified")
	fmt.Printf("-------------------------------------------------------------------------------\n")
	for _, value := range list {
		if strings.Compare(value.GetType(), "Anime") == 0 {
			val := strings.Split(value.LongOutput(), " ")
			fmt.Printf("%-20s%-10s%-10s%-15s%-15s\n", val[0], val[1], val[2], val[3], strings.Join(val[4:], " "))
		}
	}
	fmt.Println()
	/*
		Long listing for mangas
	*/
	fmt.Printf("\n%-20s%-10s%-10s%-15s%-15s\n", "Name", "Volume", "Chapter", "Publisher", "Modified")
	fmt.Printf("-------------------------------------------------------------------------------\n")
	for _, value := range list {
		if strings.Compare(value.GetType(), "Manga") == 0 {
			val := strings.Split(value.LongOutput(), " ")
			fmt.Printf("%-20s%-10s%-10s%-15s%-15s\n", val[0], val[1], val[2], val[3], strings.Join(val[4:], " "))
		}
	}
	fmt.Println()

}
