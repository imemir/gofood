package httpext

import (
	"encoding/json"
	"github.com/imemir/gofood/pkg/jsonext"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Error struct {
	Message string `json:"message,omitempty"`
}

func SendError(res http.ResponseWriter, req *http.Request, code int, message string) {
	SendModel(res, req, code, &Error{
		Message: message,
	})
}

func SendModel(res http.ResponseWriter, req *http.Request, code int, model interface{}) {
	bytes, err := json.Marshal(model)
	if err != nil {
		log.WithError(err).Error("can not marshal model")
	}
	SendData(res, req, code, jsonext.MIME, bytes)
}

func SendData(res http.ResponseWriter, req *http.Request, code int, mime string, data []byte) {
	res.Header().Set(ContentTypeHeader, mime)
	res.Header().Set(CharsetHeader, "utf-8")
	res.WriteHeader(code)
	_, err := res.Write(data)
	if err != nil {
		log.WithError(err).Error("can not write data on response")
	}
}
