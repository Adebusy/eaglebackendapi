package dataaceess

import "time"

// -- Subscriptions table: Users enroll in courses
// CREATE TABLE subscriptions (
//     subscription_id SERIAL PRIMARY KEY,
//     user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
//     course_id INT REFERENCES courses(course_id) ON DELETE CASCADE,
//     subscription_status VARCHAR(20) CHECK (subscription_status IN ('pending', 'active', 'completed', 'cancelled')) DEFAULT 'pending',
//     created_at TIMESTAMP DEFAULT NOW()
// );

type TblSubscriptions struct {
	Id                 int       `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	UserId             int       `gorm:"column:JobId"`
	CourseId           string    `gorm:"column:CourseId"`
	SubscriptionStatus string    `gorm:"column:SubscriptionStatus"`
	TransactionId      string    `gorm:"column:TransactionId"`
	CreatedAt          time.Time `gorm:"column:CreatedAt"`
}

type Subscriptions struct {
	Id                 int       `gorm:"column:Id"`
	UserId             int       `gorm:"column:JobId"`
	CourseId           string    `gorm:"column:CourseId"`
	SubscriptionStatus string    `gorm:"column:SubscriptionStatus"`
	TransactionId      string    `gorm:"column:TransactionId"`
	CreatedAt          time.Time `gorm:"column:CreatedAt"`
}
