# reqx
[![GoDoc](https://godoc.org/github.com/imroc/req?status.svg)](https://godoc.org/github.com/imroc/req)

Golang http client
- Light weight
- Simple & Easy 


Install
=======
``` sh
go get github.com/dreamph/reqx
```


Benchmark (reqx vs go http)
=======
``` sh
go test -bench . -benchmem -count 1

Benchmark_ReqxRequests/GET-12                      47740             22771 ns/op            1768 B/op         25 allocs/op
Benchmark_ReqxRequests/POST-12                     47290             24967 ns/op            1995 B/op         33 allocs/op
Benchmark_ReqxRequests/PUT-12                      47718             24997 ns/op            1991 B/op         33 allocs/op
Benchmark_ReqxRequests/PATCH-12                    47674             24923 ns/op            1992 B/op         33 allocs/op
Benchmark_ReqxRequests/DELETE-12                   47343             24903 ns/op            2004 B/op         33 allocs/op
Benchmark_GoHttpRequests/GET-12                    36661             31616 ns/op            5792 B/op         69 allocs/op
Benchmark_GoHttpRequests/POST-12                   34401             34127 ns/op            7456 B/op         88 allocs/op
Benchmark_GoHttpRequests/PUT-12                    34412             34492 ns/op            7389 B/op         88 allocs/op
Benchmark_GoHttpRequests/PATCH-12                  34220             34656 ns/op            7494 B/op         88 allocs/op
Benchmark_GoHttpRequests/DELETE-12                 34275             34700 ns/op            7473 B/op         88 allocs/op

```

Examples
=======
``` go
package main

import (
	"bytes"
	"github.com/dreamph/reqx"

	"log"
	"os"
	"time"
)

type Data struct {
	Name string `json:"name,omitempty"`
}

type Response struct {
	Origin string `json:"origin"`
}

func main() {
	client := reqx.New(&reqx.Options{
		Timeout: time.Second * 10,
		Headers: reqx.Headers{
			reqx.HeaderAuthorization: "Bearer 123456",
		},
		//InsecureSkipVerify: true,
	})

	//POST
	result := &Response{}
	resp, err := client.Post(&reqx.Request{
		URL: "https://httpbin.org/post",
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
	println(resp.StatusCode)
	println(result.Origin)

	//POST and get raw body
	var resultBytes []byte
	resp, err = client.Post(&reqx.Request{
		URL: "https://httpbin.org/post",
		Data: &Data{
			Name: "Reqx",
		},
		Result: &resultBytes,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	println(resp.StatusCode)
	println(string(resultBytes))

	//POST with request timeout
	var resultBytes2 []byte
	resp, err = client.Post(&reqx.Request{
		URL: "https://httpbin.org/post",
		Data: &Data{
			Name: "Reqx",
		},
		Result:  &resultBytes2,
		Timeout: time.Second * 5,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	println(resp.StatusCode)
	println(string(resultBytes2))

	//UPLOAD FILES
	test1Bytes, err := os.ReadFile("example/demo.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	test2Bytes, err := os.ReadFile("example/demo.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	var resultUploadBytes []byte
	resp, err = client.Post(&reqx.Request{
		URL: "https://httpbin.org/post",
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
		Result: &resultUploadBytes,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	println(resp.StatusCode)
	println(string(resultUploadBytes))

	//GET
	result = &Response{}
	resp, err = client.Get(&reqx.Request{
		URL:    "https://httpbin.org/get",
		Result: result,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	println(resp.StatusCode)
	println(result.Origin)

	//DELETE
	result = &Response{}
	resp, err = client.Delete(&reqx.Request{
		URL: "https://httpbin.org/delete",
		Data: &Data{
			Name: "Reqx",
		},
		Result: result,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	println(resp.StatusCode)
	println(result.Origin)

	//PUT
	result = &Response{}
	resp, err = client.Put(&reqx.Request{
		URL: "https://httpbin.org/put",
		Data: &Data{
			Name: "Reqx",
		},
		Headers: reqx.Headers{
			"api-key": "123456",
		},
		Result: result,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	println(resp.StatusCode)
	println(result.Origin)
}
```