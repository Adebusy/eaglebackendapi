package dataaceess

import "time"

// -- Course schedule: Calendar for training sessions
// CREATE TABLE course_schedule (
//     schedule_id SERIAL PRIMARY KEY,
//     course_id INT REFERENCES courses(course_id) ON DELETE CASCADE,
//     session_date DATE NOT NULL,
//     available_slots INT CHECK (available_slots >= 1),
//     created_at TIMESTAMP DEFAULT NOW()
// );

type TblCoursesSchedule struct {
	Id             int       `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	CourseId       int       `gorm:"column:CourseId"`
	SessionDate    time.Time `gorm:"column:SessionDate"`
	AvailableSlots int       `gorm:"column:AvailableSlots"`
	CreatedAt      time.Time `gorm:"column:CreatedAt"`
	Status         int       `gorm:"column:Status"`
}

type CoursesSchedule struct {
	Id             int       `gorm:"column:Id"`
	CourseId       int       `gorm:"column:CourseId"`
	SessionDate    time.Time `gorm:"column:SessionDate"`
	AvailableSlots int       `gorm:"column:AvailableSlots"`
	CreatedAt      time.Time `gorm:"column:CreatedAt"`
	Status         int       `gorm:"column:Status"`
}
