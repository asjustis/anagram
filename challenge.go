package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sort"

	"crypto/md5"
	"encoding/hex"
)

/* []string Sorting */

type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

// Implemented reverse order
func (a ByLen) Less(i, j int) bool {
	return len(a[i]) > len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

/* Challenge */

// Filters out words that contains letters NOT from anagram
func FilterDictionaryLetters(anagram string, dictFile string, resultFile string) {
	bs, err := ioutil.ReadFile(dictFile)
	if err != nil {
		return
	}
	str := string(bs)
	dict := strings.Split(str, "\n")


	var _dict []string
	for _, word := range dict {
		addWord := true
		for _, char := range word {
			id := strings.Index(anagram, string(char))
			if (id == -1) {
				addWord = false
			}
		}

		if (addWord) {
			_dict = append(_dict, word)
		}
	}

	sort.Sort(ByLen(_dict))
	fmt.Println(len(_dict), _dict[:10])

	f, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := range _dict {
		f.WriteString(_dict[i] + "\n")
	}
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func CheckEntry(solution string) bool{
	//md5Easy := "e4820b45d2277f3844eac66c903e84be" // answer is "printout stout yawls"
	md5Medium := "23170acc097c24edb98fc5488ab033fe" // "ty outlaws printouts"
	//md5Hard := "665e5bcb0c20062fe8abaaf4628bb154" 

	if (GetMD5Hash(solution) == md5Medium) {
		fmt.Println("Anagram: ", solution)
		return true
	}

	return false
}

// Solution for easiest and medium md5 hash
func Try2() {
	// Filter out the dictionary
	anagram := "poultry outwits ants"
	FilterDictionaryLetters(anagram, "wordlist", "wordlistCleared")

	// Read filtered dictionary
	bs, err := ioutil.ReadFile("wordlistCleared")
	if err != nil {
		return
	}
	str := string(bs)
	dict := strings.Split(str, "\n")

	// Under the assumption that anagram solution will be out of three words, 
	// we run O(n^3) loop and check all combinations for correct md5 hash
	// It proved to be good enough for easiest and medium hashes.
	for i := 0; i < len(dict); i++ {

		progress := float64(i) / float64(len(dict)) * 100.0
		fmt.Printf("Progress: %.3f\n", progress)

		for j := 0; j < len(dict); j++ {
			for k := 0; k < len(dict); k++ {
				CheckEntry(dict[i] + " " + dict[j] + " " + dict[k])
			}
		}
	}
}

// Helper function, which 'consumes' word from availableChars letter list
// Return true if each letter was used and available in availableChars
func isWordValid(word, availableChars string) (bool, string) {
	unusedChars := availableChars

	for _, letter := range word {
		id := strings.Index(unusedChars, string(letter))

		if (id == -1) {
			return false, unusedChars
		} else {
			unusedChars = unusedChars[:id] + unusedChars[id+1:]
		}
	}

	return true, unusedChars
}

// Alternative (recursive) solution in attempt to find hardest hash.
func SolveRecursively(word, availableChars string, dict []string, anagramLen, depth int) int {

	// Let's try with the max depth of 4
	if (depth > 4) {
		return -1
	}

	// Stop if current constructed word length is bigger than anagram
	if (len(strings.Replace(word, " ", "", -1)) > anagramLen) {
		return -1
	}

	// Already tried manually 2- and 3-word combinations, let's skip them
	if (depth != 2 && depth != 3) {
		if (CheckEntry(word)) {
			return 1
		}
	}

	for _, newWord := range dict {
		valid, unusedChars := isWordValid(newWord, availableChars)
		if valid {
			if (SolveRecursively(word + " " + newWord, unusedChars, dict, anagramLen, depth+1) > 0) {
				return 1
			}
		}
	}

	return -1

}

// Main function for the recursive solution
func TryHardest() {
	// Filter out the dictionary
	anagram := "poultry outwits ants"
	FilterDictionaryLetters(anagram, "wordlist", "wordlistSorted")

	// Read filtered words
	bs, err := ioutil.ReadFile("wordlistSorted")
	if err != nil {
		return
	}
	str := string(bs)
	dict := strings.Split(str, "\n")

	// Attempt to solve for hardest md5 recursively
	anagramChars := "poultryoutwitsants"
	for i, word := range dict {
		x := float64(i) / float64(len(dict)) * 100.0
		fmt.Printf("Progress: %.3f\n", x)

		valid, unusedChars := isWordValid(word, anagramChars)
		if valid {
			SolveRecursively(word, unusedChars, dict, len(anagramChars), 1)
		}
	}
}

func main() {
	Try2()
	//TryHardest()
}

