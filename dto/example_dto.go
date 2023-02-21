package dto

var ExampleDTO struct {
    ID    string `json"-"`
    Name  string `form:"name" json:"name" binding:"required"`
    Color string `form:"color" json:"color" binding:"required,oneof=blue yellow"`
}