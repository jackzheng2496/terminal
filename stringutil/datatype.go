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

//TODO:	Should prob make these functions return default Anime/Manga values

func NewAnime(name, timestamp, episode, season, studio string) *Anime {
	return &Anime{
		Episode:  episode,
		Season:   season,
		Studio:   studio,
		Japanese: Japanese{Name: name, Timestamp: timestamp},
	}
}

func NewManga(name, timestamp, chapter, novel, publisher string) *Manga {
	return &Manga{
		Chapter:   chapter,
		Volume:    novel,
		Publisher: publisher,
		Japanese:  Japanese{Name: name, Timestamp: timestamp},
	}
}

func (a *Anime) FormattedOutput() string {
	return fmt.Sprintf("%s %s %s", a.Name, a.Season, a.Episode)
}

func (a *Anime) LongOutput() string {
	return fmt.Sprintf("%s %s %s %s %s", a.Name, a.Season, a.Episode, a.Studio, a.Timestamp)
}

func (a *Anime) SaveOutput() string {
	return fmt.Sprintf("%s %s %s %s %s", a.Name, a.Season, a.Episode, a.Studio, a.Timestamp)
}

func (m *Manga) FormattedOutput() string {
	return fmt.Sprintf("%s %s %s", m.Name, m.Volume, m.Chapter)
}

func (m *Manga) LongOutput() string {
	return fmt.Sprintf("%s %s %s %s %s", m.Name, m.Volume, m.Chapter, m.Publisher, m.Timestamp)
}

func (m *Manga) SaveOutput() string {
	return fmt.Sprintf("%s %s %s %s %s", m.Name, m.Volume, m.Chapter, m.Publisher, m.Timestamp)
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

func (a *Anime) UpdateProducer(value string) {
	a.Studio = value
}

func (m *Manga) UpdateProducer(value string) {
	m.Publisher = value
}

func (a *Anime) GetType() string {
	return "Anime"
}

func (m *Manga) GetType() string {
	return "Manga"
}
