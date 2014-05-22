package api

import "encoding/json"

type Json map[string]interface{}

func (r Json) String() string {

	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(b)

}
