package stringutil

import (
	"fmt"
)

type Japan struct {
	name, timestamp string
}

type Anime struct {
	Japan
	episode, season string
}

type Manga struct {
	Japan
	chapter, novel string
}

func NewAnime(name, timestamp, episode, season string) *Anime {
	return &Anime{
		episode: episode,
		season:  season,
		Japan: Japan{
			name:      name,
			timestamp: timestamp,
		},
	}
}

func (a Anime) FormattedOutput() {
	fmt.Printf("%s Season %s Episode %s Last Modified: %s",
		a.name, a.season, a.episode, a.timestamp)
}

func (a Anime) UpdateTimeStamp(timestamp string) {
	a.timestamp = timestamp
}

func (a Anime) UpdateValue(value string) {
	a.episode = value
}

func (a Anime) UpdateSubVal(value string) {
	a.season = value
}

func NewManga(name, timestamp, chapter, novel string) *Manga {
	return &Manga{
		chapter: chapter,
		novel:   novel,
		Japan: Japan{
			name:      name,
			timestamp: timestamp,
		},
	}
}

func (m Manga) FormattedOutput() {
	fmt.Printf("%s Novel %s Chapter %s Last Modified: %s",
		m.name, m.novel, m.chapter, m.timestamp)
}

func (m Manga) UpdateTimeStamp(timestamp string) {
	m.timestamp = timestamp
}

func (m Manga) UpdateValue(value string) {
	m.chapter = value
}

func (m Manga) UpdateSubVal(value string) {
	m.novel = value
}
