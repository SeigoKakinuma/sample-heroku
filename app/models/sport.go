package models

type Sport struct {
    Id  uint64 `gorm:"primary_key" json:"id"`
    Sport string `sql:"size:255" json:"sport"`
}