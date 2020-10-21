package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
var verb_list []string
var nouns_list []string

func get_html_from_file(filename string) string {
	file_data, err := ioutil.ReadFile(filename)
	check(err)
	return string(file_data)
}

func rustle_my_jimmies(w http.ResponseWriter, r *http.Request) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	verb_index := random.Intn(len(verb_list))
	noun_index := random.Intn(len(nouns_list))
	base_html := get_html_from_file("index.html")
	t, err := template.New("test").Parse(base_html)
	check(err)
	t.ExecuteTemplate(w, "test", "Man this really " + verb_list[verb_index] + " my " + nouns_list[noun_index])
}

func init_expression_data() {
	verb_data, err := ioutil.ReadFile("verbs.txt")
	check(err)
	noun_data, err := ioutil.ReadFile("nouns.txt")
	check(err)
	verb_list = strings.Split(string(verb_data), "\n")
	nouns_list = strings.Split(string(noun_data), "\r")
}

func main() {
	init_expression_data()
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(""))))
	http.HandleFunc("/home", rustle_my_jimmies)
	log.Fatal(http.ListenAndServeTLS(":4443","server.crt" , "server.key",nil))
}

