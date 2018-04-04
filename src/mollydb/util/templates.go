package util

import (
	"strings"
)

func extractStringValueFromMap(key string, vars map[string]interface{}) string {
	//TODO This should be able to deep into the map. So far, we assume
	keys := strings.Split(key, ".")
	if len(keys) == 1 {
		val, ok := vars[key]
		if ok {
			return val.(string)
		}
		return ""
	}
	pendingVars := vars[keys[0]]

	switch val := pendingVars.(type) {
	case string:
		return val
	case map[string]interface{}:
		return extractStringValueFromMap(strings.Join(keys[1:], "."), val)
	case map[interface{}]interface{}:
		res := make(map[string]interface{})
		for k, valChild := range val {
			res[k.(string)] = valChild
		}
		return extractStringValueFromMap(strings.Join(keys[1:], "."), res)
	default:
		return ""
	}
}

//Expand function to replace variables
func Expand(s string, vars map[string]interface{}) string {
	buf := make([]byte, 0, 2*len(s))
	// ${} is all ASCII, so bytes are fine for this operation.
	i := 0
	for j := 0; j < len(s); j++ {
		if s[j] == '$' && j+1 < len(s) {
			buf = append(buf, s[i:j]...)
			name, w := getShellName(s[j+1:])
			bytes := []byte(extractStringValueFromMap(name, vars))
			buf = append(buf, bytes...)
			j += w
			i = j + 1
		}
	}
	return string(buf) + s[i:]
}

func isShellSpecialVar(c uint8) bool {
	switch c {
	case '*', '#', '$', '@', '!', '?', '-', '0', '1', '2', '3', '4', '5',
		'6', '7', '8', '9':
		return true
	}
	return false
}

func getShellName(s string) (string, int) {
	switch {
	case s[0] == '{':
		if len(s) > 2 && isShellSpecialVar(s[1]) && s[2] == '}' {
			return s[1:2], 3
		}
		// Scan to closing brace
		for i := 1; i < len(s); i++ {
			if s[i] == '}' {
				return s[1:i], i + 1
			}
		}
		return "", 1 // Bad syntax; just eat the brace.
	case isShellSpecialVar(s[0]):
		return s[0:1], 1
	}
	// Scan alphanumerics.
	var i int
	for i = 0; i < len(s) && isAlphaNum(s[i]); i++ {
	}
	return s[:i], i
}

func isAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z'
}
