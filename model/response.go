package model

type Response struct {
	StatusCode int         `json:"code,omitmplty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
