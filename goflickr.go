package goflickr

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	ENDPOINT_XMPLRPC       = "https://api.flickr.com/services/xmlrpc"
	M_FLICKR_TEST_ECHO     = "flickr.test.echo"
	M_FLICKR_PHOTOS_SEARCH = "flickr.photos.search"
	test_file              = "test.xml"
	test_file_req          = "test_req.xml"
)

type FlickrImgUrl string

type GoFlickr struct {
	ApiKey    string
	ApiSecret string
	Debug     bool
}

func New(apiKey string, apiSecret string, debug ...bool) GoFlickr {

	goFlickr := GoFlickr{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
	if len(debug) > 0 && debug[0] == true {
		goFlickr.Debug = true
	}
	return goFlickr

}

func MakeUrl(id string, server string, farm int, secret string) string {

	return "https://farm" +
		strconv.Itoa(farm) +
		".staticflickr.com/" +
		server + "/" + id + "_" + secret + "_h" + ".jpg"
}

//SearchByTags
func (g *GoFlickr) SearchByTags(tags []string) (ApiRespPhotoSearch, error) {

	if len(tags) < 1 {
		return ApiRespPhotoSearch{}, errors.New("No tags provided for search!")
	}

	concatTags := strings.Join(tags, ",")
	req := g.newRequest(M_FLICKR_PHOTOS_SEARCH)
	req.addParam("tags", concatTags)

	resp := ApiRespPhotoSearch{}
	err := req.exec(ENDPOINT_XMPLRPC, &resp)
	if err != nil {
		return ApiRespPhotoSearch{}, err
	}

	return resp, nil

}

//New returns a fully set-up Flickr Api Request object that
//can be further populated with parameters
func (g *GoFlickr) newRequest(apiMethod string) ApiRequest {

	req := ApiRequest{
		MethodName: apiMethod,
	}
	req.addParam("api_key", g.ApiKey)
	req.addParam("format", "json")
	req.addParam("nojsoncallback", "1")
	return req

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
