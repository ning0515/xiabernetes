package xiaberctl

import (
	"bytes"
	"net/http"
)

func RequestWithBody(bodyData []byte, method string) *http.Request {
	request, _ := http.NewRequest(method, "http://127.0.0.1:8000/", bytes.NewBuffer(bodyData))
	request.ContentLength = int64(len(bodyData))
	return request
}
