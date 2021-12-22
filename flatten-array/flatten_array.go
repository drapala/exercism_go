package flatten

func getNestedTypes(nested interface{}, result *[]interface{}) {
	switch nested.(type) {
	case int:
		*result = append(*result, nested)
	case []interface{}:
		for _, v := range nested.([]interface{}) { // Use type assertion to loop over []interface{}
			getNestedTypes(v, result) // Recursively call getNestedTypes
		}
	}
}

func Flatten(nested interface{}) []interface{} {
	result := make([]interface{}, 0)
	getNestedTypes(nested, &result)
	return result
}
