package httpext

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func BindModel(req *http.Request, model interface{}) (err error) {
	defer func() {
		_ = req.Body.Close()
	}()
	var obj map[string]interface{}
	if err = json.NewDecoder(req.Body).Decode(&obj); err != nil {
		return err
	}
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err = encoder.Encode(obj); err != nil {
		return err
	}
	return json.NewDecoder(&buf).Decode(model)
}
