/* Server
 */

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], ":port\n")
		os.Exit(1)
	}
	port := os.Args[1]

	http.HandleFunc("/", listFlashCards)
	fileServer := http.StripPrefix("/jscript/", http.FileServer(http.Dir("jscript")))
	http.Handle("/jscript/", fileServer)
	fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
	http.Handle("/html/", fileServer)
	fileServer = http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
	http.Handle("/css/", fileServer)

	http.HandleFunc("/flashcards.html", listFlashCards)
	http.HandleFunc("/flashcardSets", manageFlashCards)

	// deliver requests to the handlers
	err := http.ListenAndServe(port, nil)
	checkError(err)
	// That's it!
}

func listFlashCards(rw http.ResponseWriter, req *http.Request) {
	flashCardsNames := ListFlashCardsNames()
	t, err := template.ParseFiles("html/listflashcards.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
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
		set = "common_words"
		order = "Random"
		action = "Show cards in set"
		half = "Random"
	}

	cardname := "flashcardSets/" + set

	fmt.Println("set chosen is", set)
	fmt.Println("order is", order)
	fmt.Println("action is", action)
	fmt.Println("half is", half)

	fmt.Println("cardname", cardname, "action", action)

	if action == "Show cards in set" {
		showFlashCards(rw, cardname, order, half)
	} else if action == "List words in set" {
		listWords(rw, cardname)
	} else {
		fmt.Println("unknown action")
	}
}

func showFlashCards(rw http.ResponseWriter, cardname, order, half string) {
	fmt.Println("Loading card name", cardname)
	cards := new(FlashCards)
	LoadJSON(cardname, &cards)
	if order == "Sequential" {
		cards.CardOrder = "SEQUENTIAL"
	} else {
		cards.CardOrder = "RANDOM"
	}
	fmt.Println("half is", half)
	if half == "Random" {
		cards.ShowHalf = "RANDOM_HALF"
	} else if half == "English" {
		cards.ShowHalf = "ENGLISH_HALF"
	} else {
		cards.ShowHalf = "CHINESE_HALF"
	}
	fmt.Println("loaded cards", len(cards.Cards))
	fmt.Println("Card name", cards.Name)

	t := template.New("showflashcards.html")
	t = t.Funcs(template.FuncMap{"pinyin": PinyinFormatter})
	t, err := t.ParseFiles("html/showflashcards.html")
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(rw, cards)
	if err != nil {
		fmt.Println("Execute error " + err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func listWords(rw http.ResponseWriter, cardname string) {
	fmt.Println("Loading card name", cardname)
	cards := new(FlashCards)
	LoadJSON(cardname, cards)
	fmt.Println("loaded cards", len(cards.Cards))
	fmt.Println("Card name", cards.Name)

	t := template.New("listwords.html")

	t = t.Funcs(template.FuncMap{"pinyin": PinyinFormatter})
	t, err := t.ParseFiles("html/listwords.html")

	if err != nil {
		fmt.Println("Parse error " + err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(rw, cards)
	if err != nil {
		fmt.Println("Execute error " + err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
