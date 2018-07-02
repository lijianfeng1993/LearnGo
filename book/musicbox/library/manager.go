package library

import "errors"

type Music struct {
	Id string
	Name string
	Artist string
	Source string
	Type string
}

type MusicManager struct {
	musics []Music
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *Music, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) (music *Music, err error) {
	if len(m.musics) == 0 {
		return nil, nil
	}
	for _,music := range m.musics {
		if name == music.Name {
			return &music, nil
		}
	}
	return nil, nil
}

func (m *MusicManager) Add(music *Music) {
	m.musics = append(m.musics, *music)
}


func (m *MusicManager) Remove(name string) (musiclist []Music){
	index := 0
	endIndex := len(m.musics) - 1

	deleteMusic := make([]Music, 0)
	newMusics := make([]Music, 0)

	for k, music := range m.musics {
		if name == music.Name {
			newMusics = append(newMusics, m.musics[index:k]...)
			deleteMusic = append(deleteMusic, music)
			index = k + 1
		} else if k ==endIndex {
			newMusics = append(newMusics, m.musics[index:endIndex+1]...)
		}
	}

	m.musics = newMusics

	return deleteMusic
}

