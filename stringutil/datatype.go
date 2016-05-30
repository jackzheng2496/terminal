package stringutil

import (
	"fmt"
)

type Japan struct {
	name, timestamp string
}

type Anime struct {
	Japan
	episode, season int
}

type Manga struct {
	Japan
	chapter, novel int
}

func NewAnime(name string, timestamp string, episode, season int) *Anime {
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
	fmt.Printf("%s Season %d Episode %d Last Modified: %s",
		a.name, a.season, a.episode, a.timestamp)
}

func (a Anime) UpdateTimeStamp(timestamp string) {
	a.timestamp = timestamp
}

func (a Anime) UpdateValue(value int) {
	a.episode = value
}

func (a Anime) UpdateSubVal(value int) {
	a.season = value
}

func NewManga(name string, timestamp string, chapter, novel int) *Manga {
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
	fmt.Printf("%s Season %d Chapter %d Last Modified: %s",
		m.name, m.novel, m.chapter, m.timestamp)
}

func (m Manga) UpdateTimeStamp(timestamp string) {
	m.timestamp = timestamp
}

func (m Manga) UpdateValue(value int) {
	m.chapter = value
}

func (m Manga) UpdateSubVal(value int) {
	m.novel = value
}
