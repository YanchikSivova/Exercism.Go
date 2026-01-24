package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string {"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
    colors := Colors()
    res := -1
    for index, value := range colors{
        if value == color{
            res = index
        }
    }
    return res
}
