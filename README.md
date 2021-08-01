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

Examples
=======


``` go

import "reqx"

type Data struct {
	Name string `json:"name,omitempty"`
}

type Response struct {
	Status string `json:"status"`
}

func main() {
	client := reqx.New()

	//POST
	resx, err := client.Post(&reqx.Request{
		URL: "http://localhost:8080/products",
		Body: reqx.JSON(&Data{
			Name: "Reqx",
		}),
	})
	if err != nil {
		panic(err)
	}
	println(resx.StatusCode)

	result := &Response{}
	err = resx.ToJSON(result)
	if err != nil {
		panic(err)
	}
	println(result.Status)

	//GET
	resx, err = client.Get(&reqx.Request{
		URL: "http://localhost:8080/products",
	})
	if err != nil {
		panic(err)
	}
	println(resx.StatusCode)

	result = &Response{}
	err = resx.ToJSON(result)
	if err != nil {
		panic(err)
	}
	println(result.Status)

	//DELETE
	resx, err = client.Delete(&reqx.Request{
		URL: "http://localhost:8080/products",
		Body: reqx.JSON(&Data{
			Name: "Reqx",
		}),
	})
	if err != nil {
		panic(err)
	}
	println(resx.StatusCode)

	result = &Response{}
	err = resx.ToJSON(result)
	if err != nil {
		panic(err)
	}
	println(result.Status)

	//PUT
	resx, err = client.Put(&reqx.Request{
		URL: "http://localhost:8080/products",
		Body: reqx.JSON(&Data{
			Name: "Reqx",
		}),
		Headers: map[string]string{
			"api-key": "123456",
		},
	})
	if err != nil {
		panic(err)
	}
	println(resx.StatusCode)

	result = &Response{}
	err = resx.ToJSON(result)
	if err != nil {
		panic(err)
	}
	println(result.Status)
}
```