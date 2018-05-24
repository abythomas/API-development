package main 

type Article struct {
	Id int `json:"Id"` 
    Title string `json:"Title"`
    Desc string `json:"Desc"`
    Content string `json:"Content"`
}

var Articles []Article