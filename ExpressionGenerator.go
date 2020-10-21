package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type ExpressionGenerator struct {
	verb_list []string
	nouns_list []string
}


func (e *ExpressionGenerator) init_expression_data() {
	verb_data, err := ioutil.ReadFile("verbs.txt")
	check(err)
	noun_data, err := ioutil.ReadFile("nouns.txt")
	check(err)
	e.verb_list = strings.Split(string(verb_data), "\n")
	e.nouns_list = strings.Split(string(noun_data), "\r")
}

func (e *ExpressionGenerator) generateExpression() string{
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	verb_index := random.Intn(len(e.verb_list))
	noun_index := random.Intn(len(e.nouns_list))
	return "Man this really " + e.verb_list[verb_index] + " my " + e.nouns_list[noun_index]
}
