package lexer

import (
    "strings"
    "fmt"
)

func TokeniseSelect(query string) []string {

	tokens := []string{}
	queryString := strings.Split(query, " ")

	for i, word := range queryString {

		if stringInSlice(word, token["keyword"]) {
			tokens = append(tokens, []string{"keyword", word}...)
        } else if i > 0 && tokens[i] == "SELECT" || i > 0 && tokens[i] == "FROM" {
			tokens = append(tokens, []string{"table_reference", word}...)
		} else {
			//TODO once we're happy with the structure as a whole, turn this to a panic
			tokens = append(tokens, []string{"???", word}...)
		}
	}

    fmt.Println("")
    return tokens
}
