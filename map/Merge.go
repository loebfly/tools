package m

type Merge struct{}

func (receiver *Merge) Merge(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, map_ := range maps {
		for key, value := range map_ {
			result[key] = value
		}
	}
	return result
}
