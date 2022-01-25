package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
// go run dictionary.go - will load cedict_ts.u8 and render
func main() {
        // called only when debugging this package
        dictionaryPath := "cedict_ts.u8"
        d := new(Dictionary)
        d.loadDictionary(dictionaryPath)

        // check it loaded okay
        goodD := d.LookupEnglish("good")
        fmt.Println(goodD.String())
}
*/

type DictionaryEntry struct {
	Traditional  string
	Simplified   string
	Pinyin       string
	Translations []string
}

func (de DictionaryEntry) String() string {
	str := de.Traditional + de.Simplified + de.Pinyin
	for _, t := range de.Translations {
		str = str + "\n    " + t
	}
	return str
}

type Dictionary struct {
	Entries []*DictionaryEntry
}

func Init(inital_len int) *Dictionary {
	d := new(Dictionary)
	d.Entries = make([]*DictionaryEntry, inital_len)
	return d
}

func (d *Dictionary) Len() int {
	return len(d.Entries)
}

func (d *Dictionary) At(i int) *DictionaryEntry {
	return d.Entries[i]
}

func (d *Dictionary) Set(i int, x *DictionaryEntry) {
	d.Entries[i] = x
}

func (d *Dictionary) Push(x *DictionaryEntry) {
	d.Entries = append(d.Entries, x)
}

func (d *Dictionary) String() string {
	str := ""
	for n := 0; n < d.Len(); n++ {
		de := d.At(n)
		str += de.String() + "\n"
	}
	return str
}

func (d *Dictionary) LookupPinyin(py string) *Dictionary {
	newD := new(Dictionary)
	for n := 0; n < d.Len(); n++ {
		de := d.At(n)
		if de.Pinyin == py {
			newD.Push(de)
		}
	}
	return newD
}

func (d *Dictionary) LookupEnglish(eng string) *Dictionary {
	newD := new(Dictionary)
	for n := 0; n < d.Len(); n++ {
		de := d.At(n)
		for _, e := range de.Translations {
			if e == eng {
				newD.Push(de)
			}
		}
	}
	return newD
}

func (d *Dictionary) loadDictionary(path string) {

	f, err := os.Open(path)
	r := bufio.NewReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		line, err := r.ReadString('\n')
		if line[0] == '#' {
			continue
		}
		if err != nil {
			return
		}

		trad, simp, pinyin, translations := parseDictEntry(line)

		de := DictionaryEntry{
			Traditional:  trad,
			Simplified:   simp,
			Pinyin:       pinyin,
			Translations: translations}

		d.Push(&de)
	}
}

func parseDictEntry(line string) (string, string, string, []string) {
	// format is
	//    trad simp [pinyin] /trans/trans/.../
	tradEnd := strings.Index(line, " ")
	trad := line[0:tradEnd]
	line = strings.TrimSpace(line[tradEnd:])

	simpEnd := strings.Index(line, " ")
	simp := line[0:simpEnd]
	line = strings.TrimSpace(line[simpEnd:])

	pinyinEnd := strings.Index(line, "]")
	pinyin := line[1:pinyinEnd]
	line = strings.TrimSpace(line[pinyinEnd+1:])

	translations := strings.Split(line, "/")
	// includes empty at start and end, so
	translations = translations[1 : len(translations)-1]

	return trad, simp, pinyin, translations
}
