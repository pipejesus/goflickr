package goflickr

import (
	"encoding/xml"
)

type ApiRequest struct {
	XMLName    xml.Name        `xml:"methodCall"`
	MethodName string          `xml:"methodName"`
	Params     ApiRequestParam `xml:"params>param"`
}

type ApiRequestParam struct {
	XMLName xml.Name `xml:"param"`
	Value   ApiRequestValue
}

type ApiRequestValue struct {
	XMLName xml.Name `xml:"value"`
	Struct  ApiRequestStruct
}

type ApiRequestStruct struct {
	XMLName xml.Name `xml:"struct"`
	Members []ApiRequestMember
}

type ApiRequestMember struct {
	XMLName xml.Name `xml:"member"`
	Name    string   `xml:"name"`
	Value   ApiRequestMemberValue
}

type ApiRequestMemberValue struct {
	XMLName xml.Name `xml:"value"`
	String  string   `xml:"string"`
}
