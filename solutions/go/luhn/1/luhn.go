package luhn
import (
    "strings"
    "unicode"
)
func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
	total := 0
    runes := []rune(id)
    if len(runes) < 2{
        return false
    }
    check := runes[len(runes)-1]
    runes = runes[:len(runes)-1]
    lenWord := len(runes)
    for i := lenWord-1; i >= 0; i--{
        if !unicode.IsDigit(runes[i]){return false}
        num := int(runes[i]-'0')
        if (lenWord-1-i) % 2 == 0{
            num*=2
            if num > 9{
                num-=9
            }
            total+=num
        }else{
            total+=num
        }
    }
    total = (10-total%10)%10
    return total == int(check-'0')
}
