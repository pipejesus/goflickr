package goflickr

import (
	_ "encoding/json"
)

//ApiResponse is the mother of all responses. It contains the common fields.
//The Code & Message fields will only be filled if error occurs.
type ApiResp struct {
	Stat    string `json:"stat"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

//Response structure for FLICKR.PHOTOS.SEARCH
type ApiRespPhotoSearch struct {
	// Stat   string            `json:"stat"`
	Photos PhotoSearchPhotos `json:"photos"`
	ApiResp
}

type PhotoSearchPhotos struct {
	Page    int                `json:"page"`
	Pages   int                `json:"pages"`
	PerPage int                `json:"perpage"`
	Total   string             `json:"total"`
	Photo   []PhotoSearchPhoto `json:"photo,omitempty"`
}

type PhotoSearchPhoto struct {
	Id       string `json:"id"`
	Owner    string `json:"owner"`
	Secret   string `json:"secret"`
	Server   string `json:"server"`
	Farm     int    `json:"farm"`
	Title    string `json:"title"`
	IsPublic int    `json:"ispublic"`
	IsFriend int    `json:"isfriend"`
	IsFamily int    `json:"isfamily"`
}
