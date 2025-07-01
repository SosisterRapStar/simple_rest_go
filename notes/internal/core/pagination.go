package core

type LimitOffsetPaging struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Query  string `json:"query,omitempty"`
}

type CursorPaging struct {
	NextPageToken string
	Limit         int
	Query         string `json:"query,omitempty"`
}
