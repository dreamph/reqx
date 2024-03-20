package reqx

import (
	"crypto/tls"
	gojson "github.com/goccy/go-json"
	//"github.com/goccy/go-reflect"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"net/url"
	"reflect"
	"time"
)

const (
	HeaderAuthorization = "Authorization"
	HeaderContentType   = "Content-Type"
)

var (
	HeaderContentTypeJson      = "application/json"
	HeaderContentTypeJsonBytes = []byte(HeaderContentTypeJson)

	HeaderContentTypeFormUrlEncoded      = "application/x-www-form-urlencoded"
	HeaderContentTypeFormUrlEncodedBytes = []byte(HeaderContentTypeFormUrlEncoded)
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

type FormUrlEncoded struct {
	Values *url.Values
}

type Raw struct {
	Body []byte
}

type clientOptions struct {
	timeout            time.Duration
	baseURL            string
	userAgent          string
	tlsConfig          *tls.Config
	maxConnsPerHost    int
	headers            Headers
	onBeforeRequest    OnBeforeRequest
	onRequestCompleted OnRequestCompleted
	onRequestError     OnRequestError
	jsonMarshal        func(v interface{}) ([]byte, error)
	jsonUnmarshal      func(data []byte, v interface{}) error
}

type ClientOptions func(opts *clientOptions)

func WithBaseURL(baseURL string) ClientOptions {
	return func(opts *clientOptions) {
		opts.baseURL = baseURL
	}
}

func WithTimeout(timeout time.Duration) ClientOptions {
	return func(opts *clientOptions) {
		opts.timeout = timeout
	}
}

func WithUserAgent(userAgent string) ClientOptions {
	return func(opts *clientOptions) {
		opts.userAgent = userAgent
	}
}

func WithTLSConfig(tlsConfig *tls.Config) ClientOptions {
	return func(opts *clientOptions) {
		opts.tlsConfig = tlsConfig
	}
}

func WithMaxConnsPerHost(maxConnsPerHost int) ClientOptions {
	return func(opts *clientOptions) {
		opts.maxConnsPerHost = maxConnsPerHost
	}
}

func WithHeaders(headers Headers) ClientOptions {
	return func(opts *clientOptions) {
		opts.headers = headers
	}
}

func WithOnBeforeRequest(onBeforeRequest OnBeforeRequest) ClientOptions {
	return func(opts *clientOptions) {
		opts.onBeforeRequest = onBeforeRequest
	}
}

func WithOnRequestCompleted(onRequestCompleted OnRequestCompleted) ClientOptions {
	return func(opts *clientOptions) {
		opts.onRequestCompleted = onRequestCompleted
	}
}

func WithOnRequestError(onRequestError OnRequestError) ClientOptions {
	return func(opts *clientOptions) {
		opts.onRequestError = onRequestError
	}
}

func WithJsonMarshal(jsonMarshal func(v interface{}) ([]byte, error)) ClientOptions {
	return func(opts *clientOptions) {
		opts.jsonMarshal = jsonMarshal
	}
}

func WithJsonUnmarshal(jsonUnmarshal func(data []byte, v interface{}) error) ClientOptions {
	return func(opts *clientOptions) {
		opts.jsonUnmarshal = jsonUnmarshal
	}
}

type FormData map[string]string

func WithFileParams(files ...FileParam) *[]FileParam {
	return &files
}

func WithFileParam(name string, fileName string, reader io.Reader) FileParam {
	return FileParam{
		Name:     name,
		FileName: fileName,
		Reader:   reader,
	}
}

type Headers map[string]string

type RequestInfo struct {
	*fasthttp.Request
}

type ResponseInfo struct {
	*fasthttp.Response
}

type OnBeforeRequest func(req *RequestInfo)
type OnRequestCompleted func(req *RequestInfo, resp *ResponseInfo)
type OnRequestError func(req *RequestInfo, resp *ResponseInfo)

type Client interface {
	Get(request *Request) (*Response, error)
	Post(request *Request) (*Response, error)
	Put(request *Request) (*Response, error)
	Delete(request *Request) (*Response, error)
	Patch(request *Request) (*Response, error)
	Head(request *Request) (*Response, error)
	Options(request *Request) (*Response, error)
}

type httpClient struct {
	client             *fasthttp.Client
	baseURL            string
	userAgent          string
	headers            Headers
	onBeforeRequest    OnBeforeRequest
	onRequestCompleted OnRequestCompleted
	onRequestError     OnRequestError
	jsonMarshal        func(v interface{}) ([]byte, error)
	jsonUnmarshal      func(data []byte, v interface{}) error
}

func New(opts ...ClientOptions) Client {
	opt := &clientOptions{
		timeout:       time.Second * 30,
		userAgent:     defaultUserAgent,
		jsonMarshal:   gojson.Marshal,
		jsonUnmarshal: gojson.Unmarshal,
	}
	if len(opts) != 0 {
		for _, f := range opts {
			f(opt)
		}
	}

	tcpDialer := fasthttp.TCPDialer{
		Concurrency:      4096,
		DNSCacheDuration: time.Hour,
	}
	maxIdleConnDuration := time.Hour * 1
	fastHttpClient := &fasthttp.Client{
		Name:                          opt.userAgent,
		ReadTimeout:                   opt.timeout,
		WriteTimeout:                  opt.timeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
		DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
		DisablePathNormalizing:        true,
		Dial:                          tcpDialer.Dial,
		MaxConnsPerHost:               opt.maxConnsPerHost,
		TLSConfig:                     opt.tlsConfig,
	}

	c := &httpClient{
		client:             fastHttpClient,
		baseURL:            opt.baseURL,
		userAgent:          opt.userAgent,
		headers:            opt.headers,
		jsonMarshal:        opt.jsonMarshal,
		jsonUnmarshal:      opt.jsonUnmarshal,
		onBeforeRequest:    opt.onBeforeRequest,
		onRequestCompleted: opt.onRequestCompleted,
		onRequestError:     opt.onRequestError,
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

func (c *httpClient) Head(request *Request) (*Response, error) {
	return c.do(request, fasthttp.MethodHead)
}

func (c *httpClient) Options(request *Request) (*Response, error) {
	return c.do(request, fasthttp.MethodOptions)
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

	if c.onBeforeRequest != nil {
		c.onBeforeRequest(&RequestInfo{
			Request: req,
		})
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

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(
			&RequestInfo{
				Request: req,
			},
			&ResponseInfo{
				Response: resp,
			},
		)
	}

	if c.isUnSuccess(resp.StatusCode()) {
		if c.onRequestError != nil {
			c.onRequestError(
				&RequestInfo{
					Request: req,
				},
				&ResponseInfo{
					Response: resp,
				},
			)
		}
	}

	return &Response{
		StatusCode: resp.StatusCode(),
	}, nil
}

func (c *httpClient) getRequestURL(requestURL string) string {
	if c.baseURL != "" {
		return c.baseURL + requestURL
	}
	return requestURL
}

func (c *httpClient) initRequest(req *fasthttp.Request, resp *fasthttp.Response, request *Request, method string) error {
	requestURL := c.getRequestURL(request.URL)

	req.SetRequestURI(requestURL)
	req.Header.SetMethod(method)

	if c.headers != nil {
		for k, v := range c.headers {
			req.Header.Add(k, v)
		}
	}

	if request.Headers != nil {
		for k, v := range request.Headers {
			req.Header.Add(k, v)
		}
	}

	if request.Timeout > 0 {
		req.SetTimeout(request.Timeout)
	}

	if request.Data != nil {
		err := c.initContentTypeAndBodyRequest(req, resp, request, method)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *httpClient) initContentTypeAndBodyRequest(req *fasthttp.Request, _ *fasthttp.Response, request *Request, method string) error {
	contentType := request.Headers[HeaderContentType]
	rawBody, ok := c.getRawBody(request.Data)
	if ok {
		req.Header.SetContentType(contentType)
		req.SetBody(rawBody.Body)
		return nil
	}

	form, ok := c.getFormBody(request.Data)
	if ok {
		if form.FormData != nil || form.Files != nil {
			//bodyBuffer := bytes.NewBuffer(nil)
			//bodyWriter := multipart.NewWriter(bodyBuffer)

			//bodyBuffer := &bytes.Buffer{}
			//bodyWriter := multipart.NewWriter(bodyBuffer)

			bodyBuffer := getMultipartBufferPool()
			defer putMultipartBufferPool(bodyBuffer)

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

			if contentType == "" {
				req.Header.SetContentType(bodyWriter.FormDataContentType())
			}

			req.SetBody(bodyBuffer.Bytes())
		}
		return nil
	}

	formUrlEncoded, ok := c.getFormUrlEncodedBody(request.Data)
	if ok {
		if contentType == "" {
			req.Header.SetContentTypeBytes(HeaderContentTypeFormUrlEncodedBytes)
		}

		if formUrlEncoded.Values != nil {
			req.SetBodyString(formUrlEncoded.Values.Encode())
		}
		return nil
	}

	req.Header.SetContentTypeBytes(HeaderContentTypeJsonBytes)
	dataBytes, err := c.jsonMarshal(request.Data)
	if err != nil {
		return err
	}
	req.SetBodyRaw(dataBytes)

	return nil
}

func (c *httpClient) initResponse(request *Request, _ *fasthttp.Request, resp *fasthttp.Response) error {
	if c.isUnSuccess(resp.StatusCode()) {
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

func (c *httpClient) isUnSuccess(statusCode int) bool {
	return !c.isSuccess(statusCode)
}

func (c *httpClient) isSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}

func (c *httpClient) getRawBody(data interface{}) (*Raw, bool) {
	form, ok := data.(Raw)
	if ok {
		return &form, true
	}

	formPtr, ok := data.(*Raw)
	if ok {
		return formPtr, true
	}

	return nil, false
}

func (c *httpClient) getFormUrlEncodedBody(data interface{}) (*FormUrlEncoded, bool) {
	form, ok := data.(FormUrlEncoded)
	if ok {
		return &form, true
	}

	formPtr, ok := data.(*FormUrlEncoded)
	if ok {
		return formPtr, true
	}

	return nil, false
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
		_, err = fieldWriter.Write(toBytes(v))
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
