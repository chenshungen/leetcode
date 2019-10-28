package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	hashTable := make(map[string][]string)
	for _, str := range strs {
		sortStr := sortStringByCharacter(str)
		hashValue := md5.Sum([]byte(sortStr))
		hashValueStr := hex.EncodeToString(hashValue[:])
		hashTable[hashValueStr] = append(hashTable[hashValueStr], str)
	}

	res := make([][]string, 0, len(strs))

	for _, set := range hashTable {
		res = append(res, set)
	}

	return res
}

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
