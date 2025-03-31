package dataaceess

import "time"

type Status struct {
	Id        int       `gorm:"column:Id"`
	Name      string    `gorm:"column:Name"`
	Status    string    `gorm:"column:Status"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
}

type TblStatus struct {
	Id        int       `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:Name"`
	Status    string    `gorm:"column:Status"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
}
