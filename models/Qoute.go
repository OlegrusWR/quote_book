package models

type Quote struct{
	Id int `json:"id"`
	Author string `json:"author"`
	Quote string `json:"quote"`
}

