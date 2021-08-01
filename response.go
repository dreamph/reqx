package reqx

import (
	"encoding/json"
	"encoding/xml"
)

type Response struct {
	body       *[]byte
	StatusCode int
}

func (r *Response) IsSuccessful() bool {
	return r.StatusCode >= 200 && r.StatusCode <= 299
}

func (r *Response) ToJSON(result interface{}) error {
	return json.Unmarshal(*r.body, result)
}

func (r *Response) ToString() string {
	return string(*r.body)
}

func (r *Response) ToXML(result interface{}) error {
	return xml.Unmarshal(*r.body, result)
}
