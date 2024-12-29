package main

import (
	"bytes"
	"fmt"
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
	clientWithBaseURL := reqx.New(
		reqx.WithBaseURL("https://httpbin.org"),
		reqx.WithTimeout(10*time.Second),
		reqx.WithHeaders(reqx.Headers{
			reqx.HeaderAuthorization: "Bearer 123456",
		}),
		reqx.WithOnBeforeRequest(func(req *reqx.RequestInfo) {
			fmt.Println(req.String())
		}),
		reqx.WithOnRequestCompleted(func(req *reqx.RequestInfo, resp *reqx.ResponseInfo) {
			fmt.Println(resp.String())
		}),
		reqx.WithOnRequestError(func(req *reqx.RequestInfo, resp *reqx.ResponseInfo) {
			fmt.Println(resp.String())
		}),
	)

	//POST
	result := &Response{}
	resp, err := clientWithBaseURL.Post(&reqx.Request{
		URL: "/post",
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

	client := reqx.New(
		reqx.WithTimeout(10*time.Second),
		reqx.WithHeaders(reqx.Headers{
			reqx.HeaderAuthorization: "Bearer 123456",
		}),
	)

	//Example Api Style
	result = &Response{}
	res, err := reqx.Post().
		URL("https://httpbin.org/post").
		Data(&Data{
			Name: "Reqx",
		}).
		Headers(reqx.Headers{}).
		Result(result).
		Send(client)
	if err != nil {
		log.Fatalf(err.Error())
	}
	println(res.Headers)
	println(result.Origin)

	//POST
	result = &Response{}
	resp, err = client.Post(&reqx.Request{
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
				reqx.WithFileParam("file1", "test1.pdf", bytes.NewReader(test1Bytes)),
				reqx.WithFileParam("file2", "test2.pdf", bytes.NewReader(test2Bytes)),
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
