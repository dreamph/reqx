package reqx

import "net/url"

type reqJsonBody struct {
	Body interface{}
	ContentType string
}

type reqXmlBody struct {
	Body interface{}
	ContentType string
}

type reqRawBody struct {
	Body interface{}
	ContentType string
}

type reqFormBody struct {
	Body interface{}
	ContentType string
}

func JSON(body interface{}) reqJsonBody {
	return reqJsonBody{
		Body: body,
		ContentType: "application/json; charset=UTF-8",
	}
}

func Raw(body []byte, contextType string) reqRawBody {
	return reqRawBody{
		Body: body,
		ContentType: contextType,
	}
}

func Form(body url.Values) reqRawBody {
	return reqRawBody{
		Body: body,
		ContentType: "application/x-www-form-urlencoded",
	}
}

func XML(body interface{}) reqXmlBody {
	return reqXmlBody{
		Body: body,
		ContentType: "application/xml; charset=UTF-8",
	}
}
