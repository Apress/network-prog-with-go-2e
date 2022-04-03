package main

import (
	//	"fmt"
	"encoding/json"
	"io"
	"os"
	"sort"
)

type FlashCard struct {
	Simplified string
	English    string
	Dictionary *Dictionary
}

type FlashCards struct {
	Name      string
	CardOrder string
	ShowHalf  string
	Cards     []*FlashCard
}

func LoadJSON(r io.Reader, key any) {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(key)
	checkError(err)
}

func ListFlashCardsNames() []string {
	flashCardsDir, err := os.Open("flashcardsets")
	if err != nil {
		return nil
	}
	files, err := flashCardsDir.Readdir(-1)

	fileNames := make([]string, len(files))
	for n, f := range files {
		fileNames[n] = f.Name()
	}
	sort.Strings(fileNames)
	return fileNames
}
