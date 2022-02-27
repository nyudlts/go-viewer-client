package viewer

type PageCount struct {
	Label string `json:"label,omitempty"`
	Value uint   `json:"value,omitempty"`
}

type Metadata struct {
	PageCount PageCount `json:"page_count,omitempty"`
}

type Resource struct {
	Metadata Metadata `json:"metadata,omitempty"`
}

