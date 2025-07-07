package main

import "fmt"

func main() {
	words := []string{
		"Pride", "and", "Prejudice", "by", "Jane", "Austen", "Chapter", "1",
		"It", "is", "a", "truth", "universally", "acknowledged,", "that",
		"a", "single", "man", "in", "possession", "of", "a", "good", "fortune,",
		"must", "be", "in", "want", "of", "a", "wife.", "However", "little",
		"known", "the", "feelings", "or", "views", "of", "such", "a", "man",
		"may", "be", "on", "his", "first", "entering", "a", "neighbourhood,",
		"this", "truth", "is", "so", "well", "fixed", "in", "the", "minds",
		"of", "the", "surrounding", "families,", "that", "he", "is", "considered",
		"the", "rightful", "property", "of", "some", "one", "or", "other",
		"of", "their", "daughters.", "\"My", "dear", "Mr.", "Bennet,\"", "said",
		"his", "lady", "to", "him", "one", "day,", "\"have", "you", "heard",
		"that", "Netherfield", "Park", "is", "let", "at", "last?\"", "Mr.",
		"Bennet", "replied", "that", "he", "had", "not.", "\"But", "it", "is,\"",
		"returned", "she;", "\"for", "Mrs.", "Long", "has", "just", "been",
		"here,", "and", "she", "told", "me", "all", "about", "it.\"", "Mr.",
		"Bennet", "made", "no", "answer.", "\"Do", "you", "not", "want", "to",
		"know", "who", "has", "taken", "it?\"", "cried", "his", "wife",
		"impatiently.", "\"YOU", "want", "to", "tell", "me,", "and", "I", "have",
		"no", "objection", "to", "hearing", "it.\"", "\"Why", "then,\"", "said",
		"Mr.", "Bennet", "with", "some", "pleasure,", "\"you", "must", "go",
		"and", "see", "Mr.", "Bingley,", "when", "he", "comes", "into", "the",
		"neighbourhood.\"",
		// ... continue until ~1000 words
	}

	m := map[string]int{}
	for _, v := range words {
		m[v]++
	}

	for k, v := range m {
		if v > 1 {
			println("Word:", k, "Count:", v)
		}
	}
	fmt.Println("Total unique words:", len(m))
	fmt.Println("Total words:", len(words))

}
