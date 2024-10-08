package models

type Response struct {
	Code  int32       `json:"code"`
	Msg   string      `json:"msg"`
	Total int32       `json:"total"`
	Data  interface{} `json:"data"`
}
