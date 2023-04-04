package utils

func ConverToMultipleChoiceData(data interface{}) (int, []string) {
	tmp := data.(map[string]interface{})
	var max int
	options := make([]string, 0)
	for key, value := range tmp {
		if key == "options" {
			for _, opt := range value.([]interface{}) {
				options = append(options, opt.(string))
			}
		} else if key == "max" {
			max = InterfaceToInt(value)
		}
	}
	return max, options
}

func ConvertToSingleChoiceData(data interface{}) []string {
	tmp := data.(map[string]interface{})
	options := make([]string, 0)
	for key, value := range tmp {
		if key == "options" {
			for _, opt := range value.([]interface{}) {
				options = append(options, opt.(string))
			}
		}
	}
	return options
}
