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
	check(err)
	t.ExecuteTemplate(w, "test", expression_generator.generateExpression())
}


func main() {
	expression_generator = ExpressionGenerator{}
	expression_generator.init_expression_data()
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(""))))
	http.HandleFunc("/home", rustle_my_jimmies)
	log.Fatal(http.ListenAndServeTLS(":443","server.crt" , "server.key",nil))
}

