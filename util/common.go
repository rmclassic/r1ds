package util

func QueryParamOrDefault(params map[string][]string, key string, default_val interface{}) interface{} {
	if len(params[key]) == 0 {
		return default_val
	}

	return params[key][0]
}
