package model

type Page struct {
	PageNo  int     `json:"PageNo"`
	Size    int     `json:"PageSize"`
	Results []Video `json:"videos"`
}

func NewPage(pageNo int, size int) *Page {
	return &Page{Results: make([]Video, 0, size), PageNo: pageNo, Size: size}
}
