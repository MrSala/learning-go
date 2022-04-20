package mock_http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

// don't understand yet fully
func RequestBody(body interface{}) io.ReadCloser {
	b, _ := json.Marshal(body)
	return ioutil.NopCloser(bytes.NewReader(b))
}
