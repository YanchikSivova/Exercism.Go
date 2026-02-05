package parsinglogfiles
import "regexp"
func IsValidLine(text string) bool {
	reg, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
    return reg.MatchString(text)
}

func SplitLogLine(text string) []string {
	reg := regexp.MustCompile(`<[~=*\-]*>`)
    return reg.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	reg, _ := regexp.Compile(`(?i)".*password.*"`)
    counter := 0
    for _, line := range lines{
        if reg.MatchString(line){
            counter+=1
        }
    }
    return counter
}

func RemoveEndOfLineText(text string) string {
	reg, _ := regexp.Compile(`end-of-line\d*`)
    return reg.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	reg, _ := regexp.Compile(`User\s+([a-zA-Z]+[0-9]*)`)
    newLines := []string{}
    for _, line := range lines{
        fss := reg.FindStringSubmatch(line)
        if len(fss) > 0{
            prefix := "[USR] "
            prefix+= fss[1] + " "
            newLines = append(newLines, prefix+line)
        }else{
            newLines = append(newLines, line)
        }
    }
    return newLines
}
