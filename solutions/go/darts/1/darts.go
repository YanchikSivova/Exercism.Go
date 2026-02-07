package darts
import "math"
func Score(x, y float64) int {
	c := math.Sqrt(math.Pow(x, 2.0) + math.Pow(y, 2.0))
    if c > 10.0{
        return 0
    }
    if c > 5.0{
        return 1
    }
    if c > 1.0{
        return 5
    }
    return 10
}
