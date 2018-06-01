// Package bow is an application that performs bag of words
package bow

import (
	"log"
	"regexp"
	"sort"
	"strings"
)

var stopwords string

var r *regexp.Regexp

func init() {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	r = reg

	stopwords = strings.Join([]string{"i", "me", "my", "myself", "we", "our", "ours", "ourselves", "you", "your", "yours", "yourself", "yourselves", "he", "him", "his", "himself", "she", "her", "hers", "herself", "it", "its", "itself", "they", "them", "their", "theirs", "themselves", "what", "which", "who", "whom", "this", "that", "these", "those", "am", "is", "are", "was", "were", "be", "been", "being", "have", "has", "had", "having", "do", "does", "did", "doing", "a", "an", "the", "and", "but", "if", "or", "because", "as", "until", "while", "of", "at", "by", "for", "with", "about", "against", "between", "into", "through", "during", "before", "after", "above", "below", "to", "from", "up", "down", "in", "out", "on", "off", "over", "under", "again", "further", "then", "once", "here", "there", "when", "where", "why", "how", "all", "any", "both", "each", "few", "more", "most", "other", "some", "such", "no", "nor", "not", "only", "own", "same", "so", "than", "too", "very", "s", "t", "can", "will", "just", "don", "should", "now"}, " ")
}

// Parse will take an array of strings and return a processed string
func Parse(rows ...string) (words []string) {
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		s := r.ReplaceAllString(row, " ") // Remove special characters
		s = strings.ToLower(s)
		splitWords := strings.Split(s, " ")
		for _, w := range splitWords {
			if !strings.Contains(stopwords, w) {
				words = append(words, w)
			}
		}
	}
	return
}

// Dict holds the key and value of the word
type Dict struct {
	Key   string
	Value int
}

// Top will return the top n words based on count
func Top(words []string, n int) []Dict {
	kv := make(map[string]int)
	// Store the values in a map
	for i := 0; i < len(words); i++ {
		kv[words[i]]++
	}
	if len(kv) == 0 {
		return nil
	}

	var dict []Dict
	for k, v := range kv {
		dict = append(dict, Dict{k, v})
	}

	sort.SliceStable(dict, func(i, j int) bool {
		return dict[i].Value > dict[j].Value
	})

	if len(kv) < n {
		n = len(kv)
	}
	return dict[0:n]
}
