package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
	if n == 1{
        return 0, nil
    }
    if n < 1{
        return 0, errors.New("less then 1")
    }
    if n % 2 == 0{
        steps, err := CollatzConjecture(n/2)
        return steps+1, err
    }else{
        steps, err := CollatzConjecture((n*3)+1)
        return steps+1, err
    }
}
