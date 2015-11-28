package goflickr

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
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

//addParam adds a (name,value) parameter to the list of the
//parameters send as a query to Flickr API
func (a *ApiRequest) addParam(name string, value string) {

	newMember := ApiRequestMember{
		Name:  name,
		Value: ApiRequestMemberValue{String: value},
	}

	a.Params.Value.Struct.Members = append(a.Params.Value.Struct.Members, newMember)

}

//exec executes the request by marshalling the struct into XML
//and (hopefully) returning a response struct from unmarshalled JSON
//from Flickr API
func (a *ApiRequest) exec(endPoint string, response interface{}) error {

	marshalledRequest, err := xml.MarshalIndent(a, " ", "  ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(test_file_req, marshalledRequest, 0700)

	readyRequest := bytes.NewBuffer(marshalledRequest)
	resp, err := http.Post(endPoint, "text/xml", readyRequest)
	if err != nil {
		panic("Ay!")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ioutil.WriteFile(test_file, respBody, 0700)

	json.Unmarshal(respBody, &response)

	return nil

}
