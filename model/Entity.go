package model

type Entity struct {
	ID int64 `json:"id" gorm:"primaryKey:autoIncrement;column:Id"`
}
