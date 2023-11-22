package utils

import "encoding/json"

func JsonToMarshal(data interface{}) string {
	buf, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(buf)
}
