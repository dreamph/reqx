package reqx_test

import (
	"bytes"
	"fmt"
	"github.com/dreamph/reqx"
	"github.com/goccy/go-json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

type Data struct {
	Name string `json:"name,omitempty"`
}

type Response struct {
	Origin string `json:"origin"`
}

func Test_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, ToJsonString(Response{Origin: "reqx"}))
	}))
	defer ts.Close()

	client := reqx.New(
		reqx.WithTimeout(10*time.Second),
		reqx.WithHeaders(reqx.Headers{
			reqx.HeaderAuthorization: "Bearer 123456",
		}),
	)

	result := &Response{}
	resp, err := client.Get(&reqx.Request{
		URL: ts.URL,
		Data: &Data{
			Name: "Reqx",
		},
		Headers: reqx.Headers{
			"custom": "1",
		},
		Result: result,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Error("Post Error")
	}
	if result.Origin != "reqx" {
		t.Error("Post Error")
	}
}

func ToJsonString(obj interface{}) string {
	return string(ToJsonBytes(obj))
}

func ToJsonBytes(obj interface{}) []byte {
	data, err := json.Marshal(obj)
	if err != nil {
		var r []byte
		return r
	}
	return data
}

func Test_PostBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, ToJsonString(Response{Origin: "reqx"}))
	}))
	defer ts.Close()

	client := reqx.New(
		reqx.WithTimeout(10*time.Second),
		reqx.WithHeaders(reqx.Headers{
			reqx.HeaderAuthorization: "Bearer 123456",
		}),
	)

	result := &Response{}
	resp, err := client.Post(&reqx.Request{
		URL: ts.URL,
		Data: &Data{
			Name: "Reqx",
		},
		Headers: reqx.Headers{
			"custom": "1",
		},
		Result: result,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Error("Post Error")
	}
	if result.Origin != "reqx" {
		t.Error("Post Error")
	}
}

func Test_PostUploadFiles(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, ToJsonString(Response{Origin: "reqx"}))
	}))
	defer ts.Close()

	client := reqx.New(
		reqx.WithTimeout(10*time.Second),
		reqx.WithHeaders(reqx.Headers{
			reqx.HeaderAuthorization: "Bearer 123456",
		}),
	)

	test1Bytes, err := os.ReadFile("example/demo.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	test2Bytes, err := os.ReadFile("example/demo.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	result := &Response{}
	resp, err := client.Post(&reqx.Request{
		URL: ts.URL,
		Data: &reqx.Form{
			FormData: reqx.FormData{
				"firstName": "reqx",
			},
			Files: reqx.WithFileParams(
				reqx.FileParam{
					Name:     "file1",
					FileName: "test1.pdf",
					Reader:   bytes.NewReader(test1Bytes),
				},
				reqx.FileParam{
					Name:     "file2",
					FileName: "test2.pdf",
					Reader:   bytes.NewReader(test2Bytes),
				},
			),
		},
		Result: &result,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Error("Post Error")
	}
	if result.Origin != "reqx" {
		t.Error("Post Error")
	}
}

func Test_PostServerNotFound(t *testing.T) {
	client := reqx.New(
		reqx.WithTimeout(10*time.Second),
		reqx.WithHeaders(reqx.Headers{
			reqx.HeaderAuthorization: "Bearer 123456",
		}),
	)

	result := &Response{}
	resp, err := client.Get(&reqx.Request{
		URL: "https://httpbin.org/post_xxx",
		Data: &Data{
			Name: "Reqx",
		},
		Headers: reqx.Headers{
			"custom": "1",
		},
		Result: result,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	if resp.StatusCode != 404 {
		t.Error("Post Error")
	}
}

func Test_Put(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, ToJsonString(Response{Origin: "reqx"}))
	}))
	defer ts.Close()

	client := reqx.New(
		reqx.WithTimeout(10*time.Second),
		reqx.WithHeaders(reqx.Headers{
			reqx.HeaderAuthorization: "Bearer 123456",
		}),
	)

	result := &Response{}
	resp, err := client.Put(&reqx.Request{
		URL: ts.URL,
		Data: &Data{
			Name: "Reqx",
		},
		Headers: reqx.Headers{
			"custom": "1",
		},
		Result: result,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Error("Post Error")
	}
	if result.Origin != "reqx" {
		t.Error("Post Error")
	}
}

func Test_Patch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, ToJsonString(Response{Origin: "reqx"}))
	}))
	defer ts.Close()

	client := reqx.New(
		reqx.WithTimeout(10*time.Second),
		reqx.WithHeaders(reqx.Headers{
			reqx.HeaderAuthorization: "Bearer 123456",
		}),
	)

	result := &Response{}
	resp, err := client.Patch(&reqx.Request{
		URL: ts.URL,
		Data: &Data{
			Name: "Reqx",
		},
		Headers: reqx.Headers{
			"custom": "1",
		},
		Result: result,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Error("Post Error")
	}
	if result.Origin != "reqx" {
		t.Error("Post Error")
	}
}

func Test_Delete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, ToJsonString(Response{Origin: "reqx"}))
	}))
	defer ts.Close()

	client := reqx.New(
		reqx.WithTimeout(10*time.Second),
		reqx.WithHeaders(reqx.Headers{
			reqx.HeaderAuthorization: "Bearer 123456",
		}),
	)

	result := &Response{}
	resp, err := client.Delete(&reqx.Request{
		URL: ts.URL,
		Data: &Data{
			Name: "Reqx",
		},
		Headers: reqx.Headers{
			"custom": "1",
		},
		Result: result,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Error("Post Error")
	}
	if result.Origin != "reqx" {
		t.Error("Post Error")
	}
}

func Benchmark_ReqxRequests(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, ToJsonString(Response{Origin: "reqx"}))
	}))
	defer ts.Close()

	client := reqx.New(
		reqx.WithTimeout(10*time.Second),
		reqx.WithHeaders(reqx.Headers{
			reqx.HeaderAuthorization: "Bearer 123456",
		}),
	)

	b.Run("GET", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := &Response{}
			_, _ = client.Get(&reqx.Request{
				URL:    ts.URL,
				Result: result,
			})
		}
	})

	b.Run("POST", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := &Response{}
			_, _ = client.Post(&reqx.Request{
				URL: ts.URL,
				Data: &Data{
					Name: "Reqx",
				},
				Result: result,
			})
		}
	})

	b.Run("PUT", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := &Response{}
			_, _ = client.Put(&reqx.Request{
				URL: ts.URL,
				Data: &Data{
					Name: "Reqx",
				},
				Result: result,
			})
		}
	})

	b.Run("PATCH", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := &Response{}
			_, _ = client.Patch(&reqx.Request{
				URL: ts.URL,
				Data: &Data{
					Name: "Reqx",
				},
				Result: result,
			})
		}
	})

	b.Run("DELETE", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := &Response{}
			_, _ = client.Delete(&reqx.Request{
				URL: ts.URL,
				Data: &Data{
					Name: "Reqx",
				},
				Result: result,
			})
		}
	})
}

func Benchmark_GoHttpRequests(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, ToJsonString(Response{Origin: "reqx"}))
	}))
	defer ts.Close()

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	b.Run("GET", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
			if err != nil {
				log.Fatalln(err)
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}

			_, _ = io.ReadAll(resp.Body)
			_ = resp.Body.Close()

		}
	})

	b.Run("POST", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req, err := http.NewRequest(http.MethodPost, ts.URL, bytes.NewReader(ToJsonBytes(Data{
				Name: "Reqx",
			})))
			if err != nil {
				log.Fatalln(err)
			}
			req.Header.Add("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}

			_, _ = io.ReadAll(resp.Body)
			_ = resp.Body.Close()
		}
	})

	b.Run("PUT", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req, err := http.NewRequest(http.MethodPut, ts.URL, bytes.NewReader(ToJsonBytes(Data{
				Name: "Reqx",
			})))
			if err != nil {
				log.Fatalln(err)
			}
			req.Header.Add("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}

			_, _ = io.ReadAll(resp.Body)
			_ = resp.Body.Close()
		}
	})

	b.Run("PATCH", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req, err := http.NewRequest(http.MethodPatch, ts.URL, bytes.NewReader(ToJsonBytes(Data{
				Name: "Reqx",
			})))
			if err != nil {
				log.Fatalln(err)
			}
			req.Header.Add("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}

			_, _ = io.ReadAll(resp.Body)
			_ = resp.Body.Close()
		}
	})

	b.Run("DELETE", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req, err := http.NewRequest(http.MethodDelete, ts.URL, bytes.NewReader(ToJsonBytes(Data{
				Name: "Reqx",
			})))
			if err != nil {
				log.Fatalln(err)
			}
			req.Header.Add("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}

			_, _ = io.ReadAll(resp.Body)
			_ = resp.Body.Close()
		}
	})
}
