package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
var expression_generator ExpressionGenerator
func get_html_from_file(filename string) string {
	file_data, err := ioutil.ReadFile(filename)
	check(err)
	return string(file_data)
}

func rustle_my_jimmies(w http.ResponseWriter, r *http.Request) {
	base_html := get_html_from_file("index.html")
	t, err := template.New("test").Parse(base_html)
	query_verb := r.URL.Query().Get("verb")
	query_word := r.URL.Query().Get("noun")
	check(err)
	expressionData := expression_generator.generateExpression(query_verb, query_word)
	t.ExecuteTemplate(w, "test", expressionData )
}


func main() {
	expression_generator = ExpressionGenerator{}
	expression_generator.init_expression_data()
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(""))))
	http.HandleFunc("/", rustle_my_jimmies)
	log.Fatal(http.ListenAndServeTLS(":4443","server.crt" , "server.key",nil))
}

