package goflickr

type ApiRespPhotosGetSizes struct {
	ApiResp
	Sizes PhotosGetInfoSizes `json:"sizes"`
}

type PhotosGetInfoSizes struct {
	CanBlog     int                 `json:"canblog"`
	CanPrint    int                 `json:"canprint"`
	CanDownload int                 `json:"candownload"`
	Size        []PhotosGetInfoSize `json:"size,omitempty"`
}

type PhotosGetInfoSize struct {
	Label  string `json:"label"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Source string `json:"source"`
	Url    string `json:"url"`
	Media  string `json:"media"`
}
