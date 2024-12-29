package reqx

import (
	"context"
	"time"
)

type PostRequest struct {
	req *Request
}

func Post() *PostRequest {
	return &PostRequest{
		req: &Request{},
	}
}

func (r *PostRequest) URL(requestUrl string) *PostRequest {
	r.req.URL = requestUrl
	return r
}

func (r *PostRequest) Data(data any) *PostRequest {
	r.req.Data = data
	return r
}

func (r *PostRequest) Headers(headers Headers) *PostRequest {
	r.req.Headers = headers
	return r
}

func (r *PostRequest) Context(ctx context.Context) *PostRequest {
	r.req.Context = ctx
	return r
}

func (r *PostRequest) Result(result any) *PostRequest {
	r.req.Result = result
	return r
}

func (r *PostRequest) ErrorResult(errorResult any) *PostRequest {
	r.req.ErrorResult = errorResult
	return r
}

func (r *PostRequest) Timeout(timeout time.Duration) *PostRequest {
	r.req.Timeout = timeout
	return r
}

func (r *PostRequest) Send(client Client) (*Response, error) {
	return client.Post(r.req)
}

type GetRequest struct {
	req *Request
}

func Get() *GetRequest {
	return &GetRequest{
		req: &Request{},
	}
}

func (r *GetRequest) URL(requestUrl string) *GetRequest {
	r.req.URL = requestUrl
	return r
}

func (r *GetRequest) Headers(headers Headers) *GetRequest {
	r.req.Headers = headers
	return r
}

func (r *GetRequest) Context(ctx context.Context) *GetRequest {
	r.req.Context = ctx
	return r
}

func (r *GetRequest) Result(result any) *GetRequest {
	r.req.Result = result
	return r
}

func (r *GetRequest) ErrorResult(errorResult any) *GetRequest {
	r.req.ErrorResult = errorResult
	return r
}

func (r *GetRequest) Timeout(timeout time.Duration) *GetRequest {
	r.req.Timeout = timeout
	return r
}

func (r *GetRequest) Send(client Client) (*Response, error) {
	return client.Get(r.req)
}

type DeleteRequest struct {
	req *Request
}

func Delete() *DeleteRequest {
	return &DeleteRequest{
		req: &Request{},
	}
}

func (r *DeleteRequest) URL(requestUrl string) *DeleteRequest {
	r.req.URL = requestUrl
	return r
}

func (r *DeleteRequest) Data(data any) *DeleteRequest {
	r.req.Data = data
	return r
}

func (r *DeleteRequest) Headers(headers Headers) *DeleteRequest {
	r.req.Headers = headers
	return r
}

func (r *DeleteRequest) Context(ctx context.Context) *DeleteRequest {
	r.req.Context = ctx
	return r
}

func (r *DeleteRequest) Result(result any) *DeleteRequest {
	r.req.Result = result
	return r
}

func (r *DeleteRequest) ErrorResult(errorResult any) *DeleteRequest {
	r.req.ErrorResult = errorResult
	return r
}

func (r *DeleteRequest) Timeout(timeout time.Duration) *DeleteRequest {
	r.req.Timeout = timeout
	return r
}

func (r *DeleteRequest) Send(client Client) (*Response, error) {
	return client.Delete(r.req)
}

type PutRequest struct {
	req *Request
}

func Put() *PutRequest {
	return &PutRequest{
		req: &Request{},
	}
}

func (r *PutRequest) URL(requestUrl string) *PutRequest {
	r.req.URL = requestUrl
	return r
}

func (r *PutRequest) Data(data any) *PutRequest {
	r.req.Data = data
	return r
}

func (r *PutRequest) Headers(headers Headers) *PutRequest {
	r.req.Headers = headers
	return r
}

func (r *PutRequest) Context(ctx context.Context) *PutRequest {
	r.req.Context = ctx
	return r
}

func (r *PutRequest) Result(result any) *PutRequest {
	r.req.Result = result
	return r
}

func (r *PutRequest) ErrorResult(errorResult any) *PutRequest {
	r.req.ErrorResult = errorResult
	return r
}

func (r *PutRequest) Timeout(timeout time.Duration) *PutRequest {
	r.req.Timeout = timeout
	return r
}

func (r *PutRequest) Send(client Client) (*Response, error) {
	return client.Put(r.req)
}

type PatchRequest struct {
	req *Request
}

func Patch() *PatchRequest {
	return &PatchRequest{
		req: &Request{},
	}
}

func (r *PatchRequest) URL(requestUrl string) *PatchRequest {
	r.req.URL = requestUrl
	return r
}

func (r *PatchRequest) Data(data any) *PatchRequest {
	r.req.Data = data
	return r
}

func (r *PatchRequest) Headers(headers Headers) *PatchRequest {
	r.req.Headers = headers
	return r
}

func (r *PatchRequest) Context(ctx context.Context) *PatchRequest {
	r.req.Context = ctx
	return r
}

func (r *PatchRequest) Result(result any) *PatchRequest {
	r.req.Result = result
	return r
}

func (r *PatchRequest) ErrorResult(errorResult any) *PatchRequest {
	r.req.ErrorResult = errorResult
	return r
}

func (r *PatchRequest) Timeout(timeout time.Duration) *PatchRequest {
	r.req.Timeout = timeout
	return r
}

func (r *PatchRequest) Send(client Client) (*Response, error) {
	return client.Patch(r.req)
}

type HeadRequest struct {
	req *Request
}

func Head() *HeadRequest {
	return &HeadRequest{
		req: &Request{},
	}
}

func (r *HeadRequest) URL(requestUrl string) *HeadRequest {
	r.req.URL = requestUrl
	return r
}

func (r *HeadRequest) Headers(headers Headers) *HeadRequest {
	r.req.Headers = headers
	return r
}

func (r *HeadRequest) Context(ctx context.Context) *HeadRequest {
	r.req.Context = ctx
	return r
}

func (r *HeadRequest) Send(client Client) (*Response, error) {
	return client.Head(r.req)
}

type OptionsRequest struct {
	req *Request
}

func Options() *OptionsRequest {
	return &OptionsRequest{
		req: &Request{},
	}
}

func (r *OptionsRequest) URL(requestUrl string) *OptionsRequest {
	r.req.URL = requestUrl
	return r
}

func (r *OptionsRequest) Data(data any) *OptionsRequest {
	r.req.Data = data
	return r
}

func (r *OptionsRequest) Headers(headers Headers) *OptionsRequest {
	r.req.Headers = headers
	return r
}

func (r *OptionsRequest) Context(ctx context.Context) *OptionsRequest {
	r.req.Context = ctx
	return r
}

func (r *OptionsRequest) Result(result any) *OptionsRequest {
	r.req.Result = result
	return r
}

func (r *OptionsRequest) ErrorResult(errorResult any) *OptionsRequest {
	r.req.ErrorResult = errorResult
	return r
}

func (r *OptionsRequest) Timeout(timeout time.Duration) *OptionsRequest {
	r.req.Timeout = timeout
	return r
}

func (r *OptionsRequest) Send(client Client) (*Response, error) {
	return client.Options(r.req)
}
