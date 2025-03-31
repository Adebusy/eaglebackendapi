package dataaceess

import "time"

// -- Courses table: Training programs offered by each center
// CREATE TABLE courses (
//     course_id SERIAL PRIMARY KEY,
//     center_id INT REFERENCES training_centers(center_id) ON DELETE CASCADE,
//     course_name VARCHAR(255) NOT NULL,
//     description TEXT,
//     duration_weeks INT CHECK (duration_weeks > 0),
//     price DECIMAL(10,2) CHECK (price >= 0),
//     created_at TIMESTAMP DEFAULT NOW()
// );

type TblCourses struct {
	Id            int       `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	CenterId      int       `gorm:"column:CenterId"`
	CourseId      int       `gorm:"column:CourseId"`
	CourseName    string    `gorm:"column:CourseName"`
	Description   string    `gorm:"column:Description"`
	DurationWeeks int       `gorm:"column:DurationWeeks"`
	Price         string    `gorm:"column:Price"`
	CreatedAt     time.Time `gorm:"column:CreatedAt"`
	Status        int       `gorm:"column:Status"`
}

type Courses struct {
	Id            int       `gorm:"column:Id"`
	CenterId      int       `gorm:"column:CenterId"`
	CourseId      int       `gorm:"column:CourseId"`
	CourseName    string    `gorm:"column:CourseName"`
	Description   string    `gorm:"column:Description"`
	DurationWeeks int       `gorm:"column:DurationWeeks"`
	Price         string    `gorm:"column:Price"`
	CreatedAt     time.Time `gorm:"column:CreatedAt"`
	Status        int       `gorm:"column:Status"`
}
