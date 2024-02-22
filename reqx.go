package reqx

import (
	"bytes"
	"crypto/tls"
	gojson "github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"reflect"
	"time"
)

const (
	HeaderAuthorization = "Authorization"
	ContentType         = "Content-Type"
)

var (
	headerContentTypeJson = []byte("application/json")
)

const (
	defaultUserAgent = "reqx-http-client"
)

type Request struct {
	URL         string
	Data        interface{}
	Headers     Headers
	Result      interface{}
	ErrorResult interface{}
	Timeout     time.Duration
}

type Response struct {
	StatusCode int
}

type FileParam struct {
	Name     string
	FileName string
	Reader   io.Reader
}

type Form struct {
	FormData FormData
	Files    *[]FileParam
}

type Options struct {
	Timeout            time.Duration
	UserAgent          string
	InsecureSkipVerify bool
	MaxConnsPerHost    int
	Headers            map[string]string
	JsonMarshal        func(v interface{}) ([]byte, error)
	JsonUnmarshal      func(data []byte, v interface{}) error
}

type FormData map[string]string

func WithFileParams(files ...FileParam) *[]FileParam {
	return &files
}

type Headers map[string]string

type Client interface {
	Get(request *Request) (*Response, error)
	Post(request *Request) (*Response, error)
	Put(request *Request) (*Response, error)
	Delete(request *Request) (*Response, error)
	Patch(request *Request) (*Response, error)
}

type httpClient struct {
	client        *fasthttp.Client
	userAgent     string
	Headers       Headers
	jsonMarshal   func(v interface{}) ([]byte, error)
	jsonUnmarshal func(data []byte, v interface{}) error
}

func New(opts ...*Options) Client {
	opt := &Options{
		Timeout:       time.Second * 30,
		UserAgent:     defaultUserAgent,
		JsonMarshal:   gojson.Marshal,
		JsonUnmarshal: gojson.Unmarshal,
		//MaxConnsPerHost: fasthttp.DefaultMaxConnsPerHost,
	}
	if len(opts) != 0 {
		userOpt := opts[0]
		if userOpt.Timeout.Milliseconds() != 0 {
			opt.Timeout = userOpt.Timeout
		}
		if userOpt.Headers != nil {
			opt.Headers = userOpt.Headers
		}
		if userOpt.UserAgent != "" {
			opt.UserAgent = userOpt.UserAgent
		}
		if userOpt.JsonMarshal != nil {
			opt.JsonMarshal = userOpt.JsonMarshal
		}
		if userOpt.JsonUnmarshal != nil {
			opt.JsonUnmarshal = userOpt.JsonUnmarshal
		}
		if userOpt.InsecureSkipVerify {
			opt.InsecureSkipVerify = userOpt.InsecureSkipVerify
		}
		if userOpt.MaxConnsPerHost != 0 {
			opt.MaxConnsPerHost = userOpt.MaxConnsPerHost
		}
	}

	tcpDialer := fasthttp.TCPDialer{
		Concurrency:      4096,
		DNSCacheDuration: time.Hour,
	}
	maxIdleConnDuration := time.Hour * 1
	fastHttpClient := &fasthttp.Client{
		Name:                          opt.UserAgent,
		ReadTimeout:                   opt.Timeout,
		WriteTimeout:                  opt.Timeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
		DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
		DisablePathNormalizing:        true,
		Dial:                          tcpDialer.Dial,
		MaxConnsPerHost:               opt.MaxConnsPerHost,
	}

	if opt.InsecureSkipVerify {
		fastHttpClient.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	c := &httpClient{
		client:        fastHttpClient,
		userAgent:     opt.UserAgent,
		Headers:       opt.Headers,
		jsonMarshal:   opt.JsonMarshal,
		jsonUnmarshal: opt.JsonUnmarshal,
	}

	return c
}

func (c *httpClient) Get(request *Request) (*Response, error) {
	return c.do(request, fasthttp.MethodGet)
}

func (c *httpClient) Post(request *Request) (*Response, error) {
	return c.do(request, fasthttp.MethodPost)
}

func (c *httpClient) Put(request *Request) (*Response, error) {
	return c.do(request, fasthttp.MethodPut)
}

func (c *httpClient) Delete(request *Request) (*Response, error) {
	return c.do(request, fasthttp.MethodDelete)
}

func (c *httpClient) Patch(request *Request) (*Response, error) {
	return c.do(request, fasthttp.MethodPatch)
}

func (c *httpClient) do(request *Request, method string) (*Response, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	err := c.initRequest(req, resp, request, method)
	if err != nil {
		return nil, err
	}

	err = c.client.Do(req, resp)
	if err != nil {
		return &Response{
			StatusCode: resp.StatusCode(),
		}, err
	}

	err = c.initResponse(request, req, resp)
	if err != nil {
		return nil, err
	}
	return &Response{
		StatusCode: resp.StatusCode(),
	}, nil
}

func (c *httpClient) initRequest(req *fasthttp.Request, _ *fasthttp.Response, request *Request, method string) error {
	req.SetRequestURI(request.URL)
	req.Header.SetMethod(method)

	if c.Headers != nil {
		for k, v := range c.Headers {
			req.Header.Add(k, v)
		}
	}

	if request.Headers != nil {
		for k, v := range request.Headers {
			req.Header.Add(k, v)
		}
	}

	if request.Data != nil {
		form, ok := c.getFormBody(request.Data)
		if ok && (form.FormData != nil || form.Files != nil) {
			bodyBuffer := &bytes.Buffer{}
			bodyWriter := multipart.NewWriter(bodyBuffer)
			defer func() {
				_ = bodyWriter.Close()
			}()

			if form.FormData != nil {
				err := writeFieldsData(bodyWriter, form.FormData)
				if err != nil {
					return err
				}
			}

			if form.Files != nil {
				err := writeFilesData(bodyWriter, form.Files)
				if err != nil {
					return err
				}
			}

			contentType := bodyWriter.FormDataContentType()
			req.Header.SetContentType(contentType)
			req.SetBody(bodyBuffer.Bytes())
		} else {
			req.Header.SetContentTypeBytes(headerContentTypeJson)
			dataBytes, err := c.jsonMarshal(request.Data)
			if err != nil {
				return err
			}
			req.SetBodyRaw(dataBytes)
		}
	}

	if request.Timeout > 0 {
		req.SetTimeout(request.Timeout)
	}

	return nil
}

func (c *httpClient) initResponse(request *Request, _ *fasthttp.Request, resp *fasthttp.Response) error {
	if !c.isSuccess(resp.StatusCode()) {
		if request.ErrorResult != nil {
			err := c.initResult(request.ErrorResult, resp)
			if err != nil {
				return err
			}
		}
		return nil
	}

	if request.Result != nil {
		err := c.initResult(request.Result, resp)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *httpClient) initResult(result interface{}, resp *fasthttp.Response) error {
	switch result.(type) {
	case *string:
		setPointerValue(result, resp.Body())
	case *[]byte:
		setPointerValue(result, resp.Body())
	case string:
	case []byte:
	default:
		body := resp.Body()
		if body != nil {
			err := c.jsonUnmarshal(body, result)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *httpClient) isSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}

func (c *httpClient) getFormBody(data interface{}) (*Form, bool) {
	form, ok := data.(Form)
	if ok {
		return &form, true
	}

	formPtr, ok := data.(*Form)
	if ok {
		return formPtr, true
	}

	return nil, false
}

func writeFieldsData(bodyWriter *multipart.Writer, fields map[string]string) error {
	for k, v := range fields {
		fieldWriter, err := bodyWriter.CreateFormField(k)
		if err != nil {
			return err
		}
		_, err = fieldWriter.Write([]byte(v))
		if err != nil {
			return err
		}
	}
	return nil
}

func writeFilesData(bodyWriter *multipart.Writer, files *[]FileParam) error {
	for _, val := range *files {
		fileWriter, err := bodyWriter.CreateFormFile(val.Name, val.FileName)
		if err != nil {
			return err
		}
		_, err = io.Copy(fileWriter, val.Reader)
		if err != nil {
			return err
		}
	}
	return nil
}

func setPointerValue(ref interface{}, data interface{}) {
	val := reflect.ValueOf(ref)
	if val.Kind() == reflect.Ptr {
		val.Elem().Set(reflect.ValueOf(data))
	}
}
