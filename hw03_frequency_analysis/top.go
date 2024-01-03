package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordCount struct {
	word  string
	count int
}

func Top10(text string) []string {
	if text == "" {
		return []string{}
	}
	words := strings.Fields(text)

	countWords := make(map[string]int, len(words))
	wordsUniq := make([]WordCount, 0, len(countWords))

	for _, word := range words {
		countWords[word]++
	}

	for key, value := range countWords {
		word := WordCount{word: key, count: value}
		wordsUniq = append(wordsUniq, word)
	}

	sort.Slice(wordsUniq, func(i, j int) bool {
		if wordsUniq[i].count == wordsUniq[j].count {
			return wordsUniq[i].word < wordsUniq[j].word
		}
		return wordsUniq[i].count > wordsUniq[j].count
	})

	limit := len(wordsUniq)

	if limit > 10 {
		limit = 10
	}

	result := make([]string, 0, limit)

	for _, elem := range wordsUniq[:limit] {
		result = append(result, elem.word)
	}

	return result
}
