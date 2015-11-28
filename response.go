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
