package models

import "time"

type User struct {
	Id       int32     `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Time     time.Time `json:"time"`
	Status   int32     `json:"status"`
}

type UserResp struct {
	Id   int32     `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}