package viewer

import (
	"encoding/json"
)

type PageCount struct {
	Label string `json:"label,omitempty"`
	Value int    `json:"value,omitempty"`
}

type Metadata struct {
	PageCount PageCount `json:"page_count,omitempty"`
}

type Resource struct {
	Metadata Metadata `json:"metadata,omitempty"`
}

func ResourceGetByNOID(noid string) (resource Resource, err error) {
	path := "/viewer/api/v1/noid/" + noid

	body, err := GetBody(path)
	if err != nil {
		return resource, err
	}

	err = json.Unmarshal(body, &resource)
	if err != nil {
		return resource, err
	}

	return resource, nil
}
