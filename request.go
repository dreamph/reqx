package reqx

import "net/url"

const (
	ContentTypeJSON           = "application/json; charset=UTF-8"
	ContentTypeXml            = "application/xml; charset=UTF-8"
	ContentTypeFormUrlEncoded = "application/x-www-form-urlencoded"
)

type reqJsonBody struct {
	Body        interface{}
	ContentType string
}

type reqXmlBody struct {
	Body        interface{}
	ContentType string
}

type reqRawBody struct {
	Body        interface{}
	ContentType string
}

type reqFormBody struct {
	Body        interface{}
	ContentType string
}

func JSON(body interface{}) reqJsonBody {
	return reqJsonBody{
		Body:        body,
		ContentType: ContentTypeJSON,
	}
}

func Raw(body []byte, contextType string) reqRawBody {
	return reqRawBody{
		Body:        body,
		ContentType: contextType,
	}
}

func Form(body url.Values) reqRawBody {
	return reqRawBody{
		Body:        body,
		ContentType: ContentTypeFormUrlEncoded,
	}
}

func XML(body interface{}) reqXmlBody {
	return reqXmlBody{
		Body:        body,
		ContentType: ContentTypeXml,
	}
}
