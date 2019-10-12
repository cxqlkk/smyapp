package util

//增删改查的Message
type OperateMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type DataStore struct {
	Total int `json:"total"`
	Datas []interface{} `json:"datas"`
	PageSize int `json:"pageSize"`
	PageNo int `json:"pageNo"`
}
type Page struct {
	PageNo int
	PageSize int
}
type PageOption func(page *Page)
func NewPage(option ...PageOption)*Page{
	page:=&Page{
		PageNo:1,
		PageSize:5,
	}
	for _, op :=range option{
		op(page)
	}
	return page

}
func WithPageNo(pageNo int) func(page *Page) {
	return func(page *Page) {
		if pageNo!=0{
			page.PageNo=pageNo
		}

	}

}
func WithPageSize (pageSize int ) func(page *Page) {

	return func(page *Page) {
		if pageSize != 0 {
			page.PageSize = pageSize
		}

	}
}