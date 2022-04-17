package helper

func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, themap := range maps {
		for key, value := range themap {
			result[key] = value
		}
	}
	return result
}
