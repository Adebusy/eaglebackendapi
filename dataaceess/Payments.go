package dataaceess

import "time"

// Payment_id SERIAL PRIMARY KEY,
//     Job_id INT REFERENCES jobs(job_id) ON DELETE CASCADE,
//     Payer_id INT REFERENCES users(user_id) ON DELETE CASCADE,
//     Officer_id INT REFERENCES security_officers(officer_id) ON DELETE CASCADE,
//     Amount DECIMAL(10,2) CHECK (amount > 0),
//     Payment_status VARCHAR(20) CHECK (payment_status IN ('pending', 'completed', 'failed')) DEFAULT 'pending',
//     Transaction_id VARCHAR(100) UNIQUE,
//     Created_at TIMESTAMP DEFAULT NOW()

type TblPayments struct {
	Id            int       `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	JobId         int       `gorm:"column:JobId"`
	PayerId       string    `gorm:"column:PayerId"`
	OfficerId     string    `gorm:"column:OfficerId"`
	Amount        string    `gorm:"column:Amount"`
	PaymentStatus string    `gorm:"column:PaymentStatus"`
	TransactionId string    `gorm:"column:TransactionId"`
	CreatedAt     time.Time `gorm:"column:CreatedAt"`
}

type Payments struct {
	Id            int       `gorm:"column:Id"`
	JobId         int       `gorm:"column:JobId"`
	PayerId       string    `gorm:"column:PayerId"`
	OfficerId     string    `gorm:"column:OfficerId"`
	Amount        string    `gorm:"column:Amount"`
	PaymentStatus string    `gorm:"column:PaymentStatus"`
	TransactionId string    `gorm:"column:TransactionId"`
	CreatedAt     time.Time `gorm:"column:CreatedAt"`
}
