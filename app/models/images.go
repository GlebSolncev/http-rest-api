package models

import "encoding/json"

type Image struct {
	Id          uint            `json:"id"`
	Application string          `json:"application"`
	Build       int             `json:"build"`
	InputBody   json.RawMessage `json:"input_body"`
	OutputBody  json.RawMessage `json:"out_body"`
	Status      string          `json:"status" gorm:"default:'in_process'"`
	CreatedAt   string          `format:"yyy-mm-dd hh:ii:ss"`
}

func (b *Image) TableName() string {
	return "images"
}
