package json

import (
	"encoding/json"
	"io"
)


func DecodeParams(r io.Reader, view interface{}) error {
	decoder := json.NewDecoder(r)
	// TODO: Enable once the frontend is refactored to not send additional fields.
	// decoder.DisallowUnknownFields()

	err := decoder.Decode(view)
	if err != nil {
		return err
	}
	return nil
}
