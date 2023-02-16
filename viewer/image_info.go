package viewer

import (
	"encoding/json"
	"net/url"
)

type ImageInfo struct {
	Context  string `json:"@context,omitempty"`
	ID       string `json:"@id,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Width    uint32 `json:"width,omitempty"`
	Height   uint32 `json:"height,omitempty"`
	Sizes    []Size `json:"sizes,omitempty"`
}

type Size struct {
	Width  uint32 `json:"width,omitempty"`
	Height uint32 `json:"height,omitempty"`
}

func ImageInfoGetByURL(rawURL string) (imageInfo ImageInfo, err error) {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return imageInfo, err
	}

	body, err := GetBody(parsedURL.Path)
	if err != nil {
		return imageInfo, err
	}

	err = json.Unmarshal(body, &imageInfo)
	if err != nil {
		return imageInfo, err
	}

	return imageInfo, nil
}
