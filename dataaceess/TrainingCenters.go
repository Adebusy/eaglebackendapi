package dataaceess

import "time"

// -- Training centers table: Security training providers
// CREATE TABLE training_centers (
//     center_id SERIAL PRIMARY KEY,
//     user_id INT UNIQUE REFERENCES users(user_id) ON DELETE CASCADE,
//     name VARCHAR(255) NOT NULL,
//     description TEXT,
//     location GEOGRAPHY(Point, 4326), -- Allows location-based searches
//     created_at TIMESTAMP DEFAULT NOW()
// );

type TblTrainingCenters struct {
	Id          int       `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	UserId      int       `gorm:"column:UserId"`
	Name        string    `gorm:"column:Name"`
	Description string    `gorm:"column:Description"`
	Location    string    `gorm:"column:location"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
}

type TrainingCenters struct {
	Id          int       `gorm:"column:Id"`
	UserId      int       `gorm:"column:UserId"`
	Name        string    `gorm:"column:Name"`
	Description string    `gorm:"column:Description"`
	Location    string    `gorm:"column:location"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
}
