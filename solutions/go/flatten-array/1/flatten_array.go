package flatten

func Flatten(nested interface{}) []interface{} {
	res := []interface{}{}

	switch nested.(type) {
	case []interface{}:
		for _, item := range nested.([]interface{}) {
			res = append(res, Flatten(item)...)
		}
	case int:
		res = append(res, nested.(int))
	}

	return res
}
