package isogram
import "strings"
func IsIsogram(word string) bool {
	isogram := map[rune]int{}
    for _, r := range strings.ToUpper(word){
        if r == ' ' || r == '-'{
            continue
        }
        if _, ok := isogram[r]; ok {return false}
        isogram[r] = 1
    }
    return true
}
