package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep [T any](collection []T, predicate func(T) bool) []T{
    if collection == nil{return nil}
    keepCollection := []T{}
    for _, value := range collection{
        if predicate(value){
            keepCollection = append(keepCollection, value)
        }
    }
    return keepCollection
}

func Discard [T any](collection []T, predicate func(T)bool) []T{
    if collection == nil{return nil}
    discardCollection := []T{}
    for _, value := range collection{
        if !predicate(value){
            discardCollection = append(discardCollection, value)
        }
    }
    return discardCollection
}
// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
