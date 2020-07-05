package slice_utils

//Subtract -
func Subtract(originalList []interface{}, subtractList []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, original := range originalList {
		exists := false
		for _, subtract := range subtractList {
			if subtract == original {
				exists = true
			}
		}
		if !exists {
			result = append(result, original)
		}
	}
	return result
}

//ElementsInBoth -
func ElementsInBoth(firstList []interface{}, secondList []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, firstItem := range firstList {
		for _, secondItem := range secondList {
			if firstItem == secondItem {
				result = append(result, firstItem)
			}
		}
	}
	return result
}

//Pop -
func Pop(array []interface{}, key int) []interface{} {
	result := make([]interface{}, 0)
	result = append(result, array[:key]...)
	result = append(result, array[key+1:]...)
	return result
}
