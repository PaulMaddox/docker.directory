package api

import "encoding/json"

// JSON provides an easy way to write JSON responses in anything
// that requires a Stringer interface, such as the fmt.Fprintf().
// This makes it very handy for writing out JSON HTTP responses.
type JSON map[string]interface{}

func (r JSON) String() string {

	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(b)

}
