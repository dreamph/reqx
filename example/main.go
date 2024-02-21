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
