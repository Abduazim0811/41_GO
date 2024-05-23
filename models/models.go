package models

type GreetingResponse struct {
	Name   string `json:"name"`
	Length int    `json:"length"`
}