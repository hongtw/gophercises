package groupsofanagrams

import (
	"fmt"
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

func charCounter(word string) string {
	values := make([]uint8, 26)
	for _, char := range word {
		values[char-'a']++
	}
	return fmt.Sprintf("%v", values)
}
func groupsOfAnagramsV2(words []string) int {
	cache := make(map[string]empty)
	rec := make(map[string]empty)
	for _, word := range words {
		if _, ok := cache[word]; ok {
			continue
		}
		cache[word] = keyExists

		countRes := charCounter(word)
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
