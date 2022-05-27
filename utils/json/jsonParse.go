package json

import "io"

type JsonParse struct {
}

func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
