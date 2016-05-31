package stringutil

import (
	"fmt"
)

type Japanese struct {
	Name, Timestamp string
}

type Anime struct {
	Japanese
	Episode, Season, Studio string
}

type Manga struct {
	Japanese
	Chapter, Volume, Publisher string
}

func NewAnime(name, timestamp, episode, season, studio string) *Anime {
	return &Anime{
		Episode:  episode,
		Season:   season,
		Studio:   studio,
		Japanese: Japanese{Name: name, Timestamp: timestamp},
	}
}

func NewManga(name, timestamp, chapter, novel string) *Manga {
	return &Manga{
		Chapter:  chapter,
		Volume:   novel,
		Japanese: Japanese{Name: name, Timestamp: timestamp},
	}
}

func (a *Anime) FormattedOutput() {
	fmt.Printf("%s Season %s Episode %s", a.Name, a.Season, a.Episode)
}

func (a *Anime) LongOutput() {
	fmt.Printf("%s Season %s Episode %s Studio %s Modified: %s",
		a.Name, a.Season, a.Episode, a.Studio, a.Timestamp)
}

func (m *Manga) FormattedOutput() {
	fmt.Printf("%s Volume %s Chapter %s", m.Name, m.Volume, m.Chapter)
}

func (m *Manga) LongOutput() {
	fmt.Printf("%s Volume %s Chapter %s Publisher %s Modified: %s",
		m.Name, m.Volume, m.Chapter, m.Publisher, m.Timestamp)
}

func (a *Anime) UpdateTimeStamp(timestamp string) {
	a.Timestamp = timestamp
}

func (m *Manga) UpdateTimeStamp(timestamp string) {
	m.Timestamp = timestamp
}

func (a *Anime) UpdateValue(value string) {
	a.Episode = value
}

func (m *Manga) UpdateValue(value string) {
	m.Chapter = value
}

func (a *Anime) UpdateSubVal(value string) {
	a.Season = value
}

func (m *Manga) UpdateSubVal(value string) {
	m.Volume = value
}
