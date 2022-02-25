package utils

import "regexp"

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func Matches(a string, pattern regexp.Regexp) bool {
    return pattern.MatchString(a)
}

// https://stackoverflow.com/questions/66643946/how-to-remove-duplicates-strings-or-int-from-slice-in-go
func RemoveDuplicateStrings(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}