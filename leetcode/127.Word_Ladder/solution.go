package _27_Word_Ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dictionary := map[string]bool{}
	for _, word := range wordList {
		dictionary[word] = true
	}

}

func transform( words []string, endWord string, wordList map[string]bool, depth int, minDepth *int) {
	lastWord := words[len(words) - 1]
	if lastWord == endWord {
		if depth < *minDepth {
			*minDepth = depth
		}
		return
	}


}
