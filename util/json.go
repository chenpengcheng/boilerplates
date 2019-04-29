package util

import "encoding/json"

func ToString(i interface{}) string {
	b, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		return "<invalid value>"
	}
	return string(b)
}
