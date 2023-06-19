package xiaberctl

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

type Server struct {
	rawUrl string
}

func New(serverUrl string) *Server {
	return &Server{
		rawUrl: serverUrl,
	}
}
func (s *Server) Verb(verb string) *Request {
	return &Request{
		verb: verb,
		s:    s,
		path: "/",
	}
}

type Request struct {
	s    *Server
	err  error
	verb string
	path string
	body []byte
	//query   labels.Query
	query   string
	timeout time.Duration
}

func (r *Request) Path(item string) *Request {
	if r.err != nil {
		return r
	}
	r.path = path.Join(r.path, item)
	return r
}

func (r *Request) Query(item string) *Request {
	if r.err != nil {
		return r
	}
	//r.query = labels.ParseQuery(item)
	r.query = item
	return r
}

func (r *Request) Body(obj []byte) *Request {
	if r.err != nil {
		return r
	}
	r.body = obj
	return r
}

func (r *Request) Do() ([]byte, error) {
	query := url.Values{}
	finalUrl := r.s.rawUrl + r.path
	if r.query != "" {
		query.Add("labels", r.query)
	}
	finalUrl += "?" + query.Encode()
	var body io.Reader
	if r.body != nil {
		body = bytes.NewBuffer(r.body)
	}
	req, _ := http.NewRequest(r.verb, finalUrl, body)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	result, _ := io.ReadAll(response.Body)
	return result, err
}
