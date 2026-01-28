package strand

func ToRNA(dna string) string {
	rna := map[rune]string{
        'G':"C",
        'C':"G",
        'T':"A",
        'A':"U",
    }
    newRNA := ""
    for _, r := range dna{
        newRNA+=rna[r]
    }
    return newRNA
}
