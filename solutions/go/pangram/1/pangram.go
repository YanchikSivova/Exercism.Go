package pangram
import "strings"
func IsPangram(input string) bool {
	sentence := strings.ToLower(input)
    alphabet := make([]bool, 26)
    for _, ch := range sentence{
        if ch >= 'a' && ch <= 'z'{
            index := ch - 'a'
            alphabet[index] = true
        }
    }
    for _, value := range alphabet{
        if !value{
            return false
        }
    }
    return true
}
