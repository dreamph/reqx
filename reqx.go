package reqx

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Reqx struct {
	client *http.Client
}

type Options struct {
	Timeout time.Duration
}

type Request struct {
	URL     string
	Body    interface{}
	Headers map[string]string
}

func New(options ...*Options) *Reqx {
	var client = &http.Client{
		Timeout: time.Second * 60,
	}
	if len(options) != 0 {
		ops := options[0]
		client.Timeout = ops.Timeout
	}
	return &Reqx{
		client: client,
	}
}

func NewClient(client *http.Client) *Reqx {
	return &Reqx{
		client: client,
	}
}

func (r *Reqx) Post(req *Request) (*Response, error) {
	return r.Do("POST", req)
}

func (r *Reqx) Put(req *Request) (*Response, error) {
	return r.Do("PUT", req)
}

func (r *Reqx) Delete(req *Request) (*Response, error) {
	return r.Do("DELETE", req)
}

func (r *Reqx) Options(req *Request) (*Response, error) {
	return r.Do("OPTIONS", req)
}

func (r *Reqx) Head(req *Request) (*Response, error) {
	return r.Do("HEAD", req)
}

func (r *Reqx) Get(req *Request) (*Response, error) {
	return r.Do("GET", req)
}

func (r *Reqx) Patch(req *Request) (*Response, error) {
	return r.Do("PATCH", req)
}

func (r *Reqx) Do(method string, req *Request) (*Response, error) {
	endpoint, _ := url.Parse(req.URL)
	httpReq := &http.Request{
		URL:    endpoint,
		Method: method,
		Header: make(http.Header),
	}

	if req.Headers != nil {
		for key, value := range req.Headers {
			httpReq.Header.Add(key, value)
		}
	}

	if req.Body != nil {
		var data []byte
		var err error
		var ok bool
		switch vv := req.Body.(type) {
		case reqJsonBody:
			setContentType(httpReq, vv.ContentType)
			data, err = json.Marshal(vv.Body)
			if err != nil {
				return nil, err
			}
		case reqXmlBody:
			setContentType(httpReq, vv.ContentType)
			data, err = json.Marshal(vv.Body)
			if err != nil {
				return nil, err
			}
		case reqRawBody:
			setContentType(httpReq, vv.ContentType)
			data, ok = vv.Body.([]byte)
			if !ok {
				return nil, errors.New("reqRawBody")
			}
		case reqFormBody:
			setContentType(httpReq, vv.ContentType)
			body, ok := vv.Body.(url.Values)
			if !ok {
				return nil, errors.New("reqFormBody")
			}
			data = []byte(body.Encode())
		}

		httpReq.Body = io.NopCloser(bytes.NewReader(data))
		httpReq.ContentLength = int64(len(data))
	}
	resp, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	response := &Response{
		StatusCode: resp.StatusCode,
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response.body = respBody
	return response, nil
}

func setContentType(req *http.Request, contentType string) {
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", contentType)
	}
}