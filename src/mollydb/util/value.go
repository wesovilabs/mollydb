package util

import "regexp"

const (
	referenceMask = `\${(a-zA-Z.)+}`
)

var re = regexp.MustCompile(referenceMask)

//HashReference check if has references
func HashReference(value string) bool {
	return re.Match([]byte(value))
}
