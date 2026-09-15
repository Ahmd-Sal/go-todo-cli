package models

type Todo struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}
