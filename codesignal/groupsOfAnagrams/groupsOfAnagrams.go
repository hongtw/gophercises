package groupsofanagrams

import (
	"sort"
	"strings"
	"sync"
)

type empty struct{}

var keyExists = empty{}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupsOfAnagrams(words []string) int {
	cache := make(map[string]empty)
	rec := make(map[string]empty)
	for _, word := range words {
		if _, ok := cache[word]; ok {
			continue
		}
		cache[word] = keyExists

		sortedWord := sortString(word)
		rec[sortedWord] = keyExists
	}
	return len(rec)
}

var globalvalues [26]uint8

func charCounter(word string) [26]uint8 {

	for i := range globalvalues {
		globalvalues[i] = 0
	}

	for _, char := range word {
		globalvalues[char-'a']++
	}
	return globalvalues
}

func groupsOfAnagramsV2(words []string) int {
	cache := make(map[string]empty)
	rec := make(map[[26]uint8]empty)
	for _, word := range words {
		if _, ok := cache[word]; ok {
			continue
		}
		cache[word] = keyExists

		countRes := charCounter(word)
		// charCounter(word)
		rec[countRes] = keyExists
	}
	return len(rec)
}

func groupsOfAnagramsV3(words []string) int {
	var wg sync.WaitGroup
	cache := make(map[string]empty)
	rec := make(map[string]empty)
	keys := make(chan string, 10000)
	done := make(chan bool)

	for _, word := range words {
		if _, ok := cache[word]; ok {
			continue
		}
		wg.Add(1)
		cache[word] = keyExists

		go func(w string, wg *sync.WaitGroup, keys chan<- string) {
			defer wg.Done()
			keys <- sortString(w)
		}(word, &wg, keys)
	}
	go func() {
		for {
			key, ok := <-keys
			if !ok {
				break
			}
			rec[key] = keyExists
		}
		done <- true
	}()
	wg.Wait()
	close(keys)
	<-done
	return len(rec)
}

func groupsOfAnagramsV4(words []string) int {
	var wg sync.WaitGroup
	cache := make(map[string]empty)
	rec := make(map[string]empty)
	keys := make(chan string, 10000)
	done := make(chan bool)

	for _, word := range words {
		if _, ok := cache[word]; ok {
			continue
		}
		wg.Add(1)
		cache[word] = keyExists

		go func(w string) {
			defer wg.Done()
			keys <- sortString(w)
		}(word)
	}
	go func() {
		for {
			key, ok := <-keys
			if !ok {
				break
			}
			rec[key] = keyExists
		}
		done <- true
	}()
	wg.Wait()
	close(keys)
	<-done
	return len(rec)
}

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}
func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortStringV2(s string) string {
	r := []rune(s) // need this, if just pass string into sortRunes(), the result won't be saved, why?
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupsOfAnagramsMaxVersion(words []string) int {
	set := make(map[string]struct{})
	for _, w := range words {
		sorted := sortStringV2(w)
		set[sorted] = struct{}{}
	}
	return len(set)
}
