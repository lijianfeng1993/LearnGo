package library

import (
	"testing"
)

func TestOpt(t *testing.T) {
	mm := NewMusicManager()

	if mm == nil {
		t.Error("NewMusicManager failed.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty.")
	}

	m0 := &Music{"1", "My heart will go on", "ljf","http://qbox.me/24501234", "mp3"}

	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed.")
	}

	mname, _ := mm.Find(m0.Name)
	if mname == nil {
		t.Error("MusicManager.Find() failed.")
	}

	deletename := mm.Remove(m0.Name)
	if deletename[0].Name != m0.Name {
		t.Error("MusicManager.Remove failed.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty.")
	}

	/*
	needname, _ := mm.Find(m0.Name)
	if needname.Name == m0.Name {
		t.Error("MusicManager.Remove failed.")
	}
	*/
}

