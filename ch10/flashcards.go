package main

import (
	//	"fmt"
	"encoding/json"
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

func LoadJSON(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}

func ListFlashCardsNames() []string {
	flashcardsDir, err := os.Open("flashcardsets")
	if err != nil {
		return nil
	}
	files, err := flashcardsDir.Readdir(-1)

	fileNames := make([]string, len(files))
	for n, f := range files {
		fileNames[n] = f.Name()
	}
	sort.Strings(fileNames)
	return fileNames
}
