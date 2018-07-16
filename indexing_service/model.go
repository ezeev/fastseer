package main

type IndexingStatusResponse struct {
	Shop            string `json:"shop"`
	Message         string `json:"message"`
	ProductsToCrawl int    `json:"productsToCrawl"`
	PageSize        int    `json:"pageSize"`
}
