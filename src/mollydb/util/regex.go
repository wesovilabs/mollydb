package util

import (
	"errors"
	"regexp"
)

//Match match rhe string and returns the values
func Match(regEx, text string) (map[string]string, error) {
	re := regexp.MustCompile(regEx)
	if re.Match([]byte(text)) {
		return extract(re, text), nil
	}
	return nil, errors.New("No matched")
}

func extract(regEx *regexp.Regexp, text string) (paramsMap map[string]string) {
	match := regEx.FindStringSubmatch(text)
	paramsMap = make(map[string]string)
	for i, name := range regEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
