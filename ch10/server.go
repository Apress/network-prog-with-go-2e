/* Server
 */

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	DefaultSet           = "common_words"
	DefaultAmount        = "Random"
	ActionShow           = "Show cards in set"
	ActionList           = "List words in set"
	ActionUnknown        = "Unknown action"
	URLFlashCardSetsPath = "flashcardSets"
	FlashCardPage        = "flashcards.html"
	ListFlashCardPage    = "list" + FlashCardPage
	ShowFlashCardPage    = "show" + FlashCardPage
	ListWordsPage        = "listwords.html"
	CardOrderSequential  = "Sequential"
	CardOrderRandom      = "Random"
)

var showHalf = map[string]string{
	"Random":  "RANDOM_HALF",
	"English": "ENGLISH_HALF",
	"Chinese": "CHINESE_HALF",
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], ":port")
	}
	port := os.Args[1]

	http.HandleFunc("/", listFlashCards)
	fileServer := http.StripPrefix("/jscript/", http.FileServer(http.Dir("jscript")))
	http.Handle("/jscript/", fileServer)
	fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
	http.Handle("/html/", fileServer)
	fileServer = http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
	http.Handle("/css/", fileServer)

	http.HandleFunc("/"+FlashCardPage, listFlashCards)
	http.HandleFunc("/"+URLFlashCardSetsPath, manageFlashCards)

	// deliver requests to the handlers
	err := http.ListenAndServe(port, nil)
	checkError(err)
}

func listFlashCards(rw http.ResponseWriter, req *http.Request) {
	flashCardsNames := ListFlashCardsNames()
	t, err := template.ParseFiles("html/" + ListFlashCardPage)

	if err != nil {
		httpErrorHandler(rw, err)
		return
	}

	t.Execute(rw, flashCardsNames)
}

/*
 * Called from listflashcards.html on form submission
 */
func manageFlashCards(rw http.ResponseWriter, req *http.Request) {
	set := req.FormValue("flashcardSets")
	order := req.FormValue("order")
	action := req.FormValue("submit")
	half := req.FormValue("half")

	//if unset
	//http://localhost:8000/flashcardSets?flashcardSets=common_words&order=Random&half=Random&submit=Show+cards+in+set
	if len(set) == 0 {
		set = DefaultSet
		order = DefaultAmount
		action = ActionShow
		half = DefaultAmount
	}

	cardname := URLFlashCardSetsPath + "/" + set

	fmt.Printf("Set %s, order %s, action %s, half %s, cardname %s\n", set, order, action, half, cardname)

	switch action {
	case ActionShow:
		showFlashCards(rw, cardname, order, half)
	case ActionList:
		listWords(rw, cardname)
	default:
		fmt.Println(ActionUnknown)
	}
}

func showFlashCards(rw http.ResponseWriter, cardname, order, half string) {
	cards := new(FlashCards)
	content, err := os.Open(cardname)
	checkError(err)
	LoadJSON(content, &cards)

	switch order {
	case CardOrderSequential:
		cards.CardOrder = "SEQUENTIAL"
	default:
		cards.CardOrder = "RANDOM"
	}

	if v, ok := showHalf[half]; ok {
		cards.ShowHalf = v
	} else {
		cards.ShowHalf = showHalf["Chinese"]
	}

	fmt.Printf("Loading card %s, half %s, loaded # of %d, card name %s\n", cardname, half, len(cards.Cards), cards.Name)

	t, err := template.New(ShowFlashCardPage).Funcs(template.FuncMap{"pinyin": PinyinFormatter}).ParseFiles("html/" + ShowFlashCardPage)

	if err != nil {
		httpErrorHandler(rw, err)
		return
	}

	err = t.Execute(rw, cards)

	if err != nil {
		httpErrorHandler(rw, err)
		return
	}
}

func listWords(rw http.ResponseWriter, cardname string) {
	cards := new(FlashCards)
	content, err := os.Open(cardname)
	checkError(err)
	LoadJSON(content, &cards)

	fmt.Printf("Loading card name %s, loaded cards %d, card name %s\n", cardname, len(cards.Cards), cards.Name)

	t, err := template.New(ListWordsPage).Funcs(template.FuncMap{"pinyin": PinyinFormatter}).ParseFiles("html/" + ListWordsPage)

	if err != nil {
		httpErrorHandler(rw, err)
		return
	}

	err = t.Execute(rw, cards)

	if err != nil {
		httpErrorHandler(rw, err)
		return
	}
}

func httpErrorHandler(rw http.ResponseWriter, err error) {
	http.Error(rw, err.Error(), http.StatusInternalServerError)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
