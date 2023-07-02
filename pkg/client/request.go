package client

import (
	"bytes"
	"fmt"
	"github.com/learnk8s/xiabernetes/pkg/api"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

func (c *Client) Verb(verb string) *Request {
	return &Request{
		verb:       verb,
		c:          c,
		path:       "/",
		pollPeriod: 1 * time.Second,
	}
}

func (c *Client) Post() *Request {
	return c.Verb("POST")
}

// Begin a PUT request.
func (c *Client) Put() *Request {
	return c.Verb("PUT")
}

// Begin a GET request.
func (c *Client) Get() *Request {
	return c.Verb("GET")
}

// Begin a DELETE request.
func (c *Client) Delete() *Request {
	return c.Verb("DELETE")
}

func (c *Client) PollFor(operationId string) *Request {
	r := c.Get().Path("operations").Path(operationId).PollPeriod(0)
	fmt.Println("url是：" + r.path)
	return r
}

type Request struct {
	c    *Client
	err  error
	verb string
	path string
	body []byte
	//query   labels.Query
	query      string
	timeout    time.Duration
	pollPeriod time.Duration
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

func (r *Request) PollPeriod(time time.Duration) *Request {
	if r.err != nil {
		return r
	}
	r.pollPeriod = time
	return r
}
func (r *Request) Do() ([]byte, error) {
	for {
		query := url.Values{}
		finalUrl := r.c.host + r.path
		if r.query != "" {
			query.Add("labels", r.query)
		}
		finalUrl += "?" + query.Encode()
		var body io.Reader
		if r.body != nil {
			body = bytes.NewBuffer(r.body)
		}
		req, _ := http.NewRequest(r.verb, finalUrl, body)
		result, err := r.c.doRequest(req)
		if err != nil {
			if statusErr, ok := err.(*StatusErr); ok {
				if statusErr.Status.Status == api.StatusWorking && r.pollPeriod != 0 {
					time.Sleep(r.pollPeriod)
					// Make a poll request
					pollOp := r.c.PollFor(statusErr.Status.Details).PollPeriod(r.pollPeriod)
					// Could also say "return r.Do()" but this way doesn't grow the callstack.
					r = pollOp
					continue
				}
			}
		}
		return result, err
	}
}
