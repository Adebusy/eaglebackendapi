package dataaceess

import "time"

// Job_id SERIAL PRIMARY KEY,
//     Requester_id INT REFERENCES users(user_id) ON DELETE CASCADE,
//     Officer_id INT REFERENCES security_officers(officer_id) ON DELETE CASCADE,
//     Job_location TEXT, -- Job site location
//     Start_time TIMESTAMP NOT NULL,
//     End_time TIMESTAMP NOT NULL,
//     Total_hours INT GENERATED ALWAYS AS (EXTRACT(EPOCH FROM (end_time - start_time)) / 3600) STORED,
//     Status VARCHAR(20) CHECK (status IN ('pending', 'confirmed', 'completed', 'cancelled')) DEFAULT 'pending',
//     Created_at TIMESTAMP DEFAULT NOW()

type TblJobs struct {
	Id          int       `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	RequesterId int       `gorm:"column:RequesterId"`
	OfficerId   string    `gorm:"column:OfficerId"`
	JobLocation string    `gorm:"column:JobLocation"`
	StartTime   string    `gorm:"column:StartTime"`
	EndTime     string    `gorm:"column:EndTime"`
	TotalHours  string    `gorm:"column:TotalHours"`
	Status      int       `gorm:"column:Status"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
}

type Jobs struct {
	Id          int       `gorm:"column:Id"`
	RequesterId int       `gorm:"column:RequesterId"`
	OfficerId   string    `gorm:"column:OfficerId"`
	JobLocation string    `gorm:"column:JobLocation"`
	StartTime   string    `gorm:"column:StartTime"`
	EndTime     string    `gorm:"column:EndTime"`
	TotalHours  string    `gorm:"column:TotalHours"`
	Status      int       `gorm:"column:Status"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
}
