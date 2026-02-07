// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob
import (
    "unicode/utf8"
    "strings"
    "regexp"
)
// Hey should have a comment documenting it.
func Hey(remark string) string {
    remark = strings.TrimSpace(remark)
	lastRune, _ := utf8.DecodeLastRuneInString(remark)
    re := regexp.MustCompile(`[a-zA-z]`)
    containsLetters := re.MatchString(remark)
    if remark == ""{
        return "Fine. Be that way!"
    }
    if lastRune == '?' && remark == strings.ToUpper(remark) && containsLetters{
        return "Calm down, I know what I'm doing!"
    }
    if lastRune == '?'{
        return "Sure."
    }
    if remark == strings.ToUpper(remark) && containsLetters{
        return "Whoa, chill out!"
    }
	return "Whatever."
}
