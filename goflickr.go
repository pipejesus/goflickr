package goflickr

import (
	"errors"
	"strconv"
	"strings"
)

const (
	ENDPOINT_XMPLRPC         = "https://api.flickr.com/services/xmlrpc"
	M_FLICKR_TEST_ECHO       = "flickr.test.echo"
	M_FLICKR_PHOTOS_SEARCH   = "flickr.photos.search"
	M_FLICKR_PHOTOS_GETSIZES = "flickr.photos.getSizes"
	test_file                = "test.json"
	test_file_req            = "test_req.xml"
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

func MakeSourceUrl(id string, server string, farm int, secret string) string {

	return "https://farm" +
		strconv.Itoa(farm) +
		".staticflickr.com/" +
		server + "/" + id + "_" + secret + "_h" + ".jpg"
}

//SearchByTags
func (g *GoFlickr) PhotosSearchByTags(tags []string) (ApiRespPhotosSearch, error) {

	if len(tags) < 1 {
		return ApiRespPhotosSearch{}, errors.New("No tags provided for search!")
	}

	concatTags := strings.Join(tags, ",")
	req := g.newRequest(M_FLICKR_PHOTOS_SEARCH)
	req.addParam("tags", concatTags)

	resp := ApiRespPhotosSearch{}
	err := req.exec(ENDPOINT_XMPLRPC, &resp)
	if err != nil {
		return ApiRespPhotosSearch{}, err
	}

	return resp, nil

}

//PhotosGetSizes
func (g *GoFlickr) PhotosGetSizes(photoId string) (ApiRespPhotosGetSizes, error) {

	if photoId == "" {
		return ApiRespPhotosGetSizes{}, errors.New("No photoID provided")
	}

	req := g.newRequest(M_FLICKR_PHOTOS_GETSIZES)
	req.addParam("photo_id", photoId)

	resp := ApiRespPhotosGetSizes{}
	err := req.exec(ENDPOINT_XMPLRPC, &resp)
	if err != nil {
		return ApiRespPhotosGetSizes{}, err
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
