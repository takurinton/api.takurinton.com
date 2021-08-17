package utils

import (
	"regexp"
	"strings"
)

func ParseContents(s string) string {
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "。  \n", "。", -1)
	s = strings.Replace(s, "。 \n", "。", -1)
	s = strings.Replace(s, "．  \n", "。", -1)
	s = strings.Replace(s, "\n", " ", -1)
	s = strings.Replace(s, "#", "", -1)
	rep := regexp.MustCompile(`\[.*\]\(https?://[\w/:%#\$&\?\(\)~\.=\+\-]+\)`)
	s = rep.ReplaceAllString(s, "")
	rep = regexp.MustCompile("```.+```")
	s = rep.ReplaceAllString(s, "")

	return s
}

func getCondition(condition string) (mental, physical int) {
	return
}
