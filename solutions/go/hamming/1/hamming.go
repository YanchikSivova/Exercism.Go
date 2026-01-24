package hamming
import (
    "unicode/utf8"
    "errors"
)
func Distance(a, b string) (int, error) {
	len_a := utf8.RuneCountInString(a)
    len_b := utf8.RuneCountInString(b)
    if len_a != len_b{
        return 0, errors.New("different length")
    }
    hammingDistance := 0
    for i:=0; i < len_a; i++{
        if a[i] != b[i]{
            hammingDistance+=1
        }
    }
    return hammingDistance, nil
}
