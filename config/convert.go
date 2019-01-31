package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	yaml "gopkg.in/yaml.v2"
)

// Convert a readcloser json or yaml to object
func Convert(r io.ReadCloser, v interface{}) error {

	// formating the readcloser to strings template vars
	b := new(bytes.Buffer)
	b.ReadFrom(r)

	// parsing json
	if json.Unmarshal(b.Bytes(), v) == nil {
		return nil
	}

	// parsing yaml
	if yaml.Unmarshal(b.Bytes(), v) == nil {
		return nil
	}

	return errors.New("Error when parsing json or yaml to object")
}
