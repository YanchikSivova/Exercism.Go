package dndcharacter
import (
    "math/rand"
    "math"
)
type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	res := int(math.Floor(float64(score - 10) / 2))
    return res
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
    dice := [4]int{
    	rand.Intn(6)+1,
        rand.Intn(6)+1,
        rand.Intn(6)+1,
        rand.Intn(6)+1,
    }
    min := 7
    total := 0
    for _, v := range dice{
        total+= v
        if v < min{
            min = v
        }
    }
    return total-min
}

// func min(a, b int) int{
//     if b < a{
//         return b
//     }
//     return a
// }

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	var char Character = Character{
        Strength: Ability(),
		Dexterity: Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom: Ability(),
		Charisma: Ability(),
    }
    char.Hitpoints = 10 + Modifier(char.Constitution)
    return char
}
