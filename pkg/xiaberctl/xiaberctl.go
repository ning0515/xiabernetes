package xiaberctl

import (
	"bytes"
	"net/http"
)

func RequestWithBody(bodyData []byte, url, method string) *http.Request {
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(bodyData))
	request.ContentLength = int64(len(bodyData))
	return request
}
