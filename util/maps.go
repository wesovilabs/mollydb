package util

//NormalizeMap function
func NormalizeMap(in map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for i, variable := range in {
		switch val := variable.(type) {
		case map[interface{}]interface{}:
			res[i] = NormalizeMap(normalizeGenericMap(val))
		case *map[string]interface{}:
			res[i] = NormalizeMap(*val)
		default:
			res[i] = val
		}
	}
	return res
}

func normalizeGenericMap(in map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, valChild := range in {
		res[k.(string)] = valChild
	}
	return res
}

//Replace function
func Replace(value interface{}, vars map[string]interface{}) interface{} {
	switch val := value.(type) {
	case string:
		return replaceString(val, vars)
	case map[string]interface{}:
		for k, valChild := range val {
			val[k] = Replace(valChild, vars)
		}
		return val
	case map[interface{}]interface{}:
		res := make(map[string]interface{})
		for k, valChild := range val {
			res[k.(string)] = Replace(valChild, vars)
		}
		return res
	case []interface{}:
		for k, valChild := range val {
			val[k] = Replace(valChild, vars)
		}
		return val
	}
	return nil
}

func replaceString(value string, vars map[string]interface{}) string {
	return Expand(value, vars)
}

func replace(value interface{}, vars map[string]interface{}) interface{} {
	switch val := value.(type) {
	case string:
		return replaceString(val, vars)
	case map[string]interface{}:
		for k, valChild := range val {
			val[k] = replace(valChild, vars)
		}
		return val
	case map[interface{}]interface{}:
		res := make(map[string]interface{})
		for k, valChild := range val {
			res[k.(string)] = replace(valChild, vars)
		}
		return res
	case []interface{}:
		for k, valChild := range val {
			val[k] = replace(valChild, vars)
		}
		return val
	}
	return nil
}

//Get function
func Get(rs map[string]interface{}, keys []string) interface{} {

	if len(keys) == 1 {
		return rs[keys[0]]
	}
	val := make(map[string]interface{})
	for _, key := range keys {
		out := rs[key]
		switch outVal := out.(type) {
		case string:
			return outVal
		case map[interface{}]interface{}:
			return Get(normalizeGenericMap(outVal), keys[1:])
		}
	}
	return val
}
