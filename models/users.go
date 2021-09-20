package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json: "name"`
	Age       int    `json: "age"`
	Sex       string `json: "sex"`
	Client_id string `json: "client_id"`
}
