package dataaceess

import "time"

// Rating_id SERIAL PRIMARY KEY,
//     Job_id INT REFERENCES jobs(job_id) ON DELETE CASCADE,
//     Officer_id INT REFERENCES security_officers(officer_id) ON DELETE CASCADE,
//     Rated_by INT REFERENCES users(user_id) ON DELETE CASCADE,
//     Rating_value INT CHECK (rating_value BETWEEN 1 AND 5) NOT NULL,
//     Review_text TEXT,
//     Created_at TIMESTAMP DEFAULT NOW()

type TblRatings struct {
	Id          int       `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	JobId       int       `gorm:"column:JobId"`
	OfficerId   string    `gorm:"column:OfficerId"`
	RatedBy     int       `gorm:"column:RatedBy"`
	RatingValue string    `gorm:"column:RatingValue"`
	ReviewText  string    `gorm:"column:ReviewText"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
}

type Ratings struct {
	Id          int       `gorm:"column:Id"`
	JobId       int       `gorm:"column:JobId"`
	OfficerId   string    `gorm:"column:OfficerId"`
	RatedBy     int       `gorm:"column:RatedBy"`
	RatingValue string    `gorm:"column:RatingValue"`
	ReviewText  string    `gorm:"column:ReviewText"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
}
