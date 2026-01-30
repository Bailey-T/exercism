package parsinglogfiles

import (
	//"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(~|\*|=|-)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _,v := range lines {
		if re.MatchString(v) {
			count ++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line(\d)+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(?i).*(User)\s+([a-z]+\d+)\s+`)
	for i,v := range lines {
		if re.MatchString(v) {
			// Really useful for seing the Submatches in the array:
			// fmt.Printf("%q\n", re.FindStringSubmatch(v))
			matches := re.FindStringSubmatch(v)
			lines[i] = "[USR] " + matches[2]+ " " + lines[i]
		}
	}
	return lines
}
