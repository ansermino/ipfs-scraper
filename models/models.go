package models

type PageInfo struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}
type PageVersion struct {
	Url       string `json:"url"`
	Title     string `json:"title"`
	Timestamp string `json:"timestamp"`
	Cid       string `json:"cid"`
}

type PageVersions []PageVersion
