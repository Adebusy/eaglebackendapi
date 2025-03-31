package dataaceess

import "time"

// Officer_id SERIAL PRIMARY KEY,
//     User_id INT UNIQUE REFERENCES users(user_id) ON DELETE CASCADE,
//     Date_of_birth DATE NOT NULL,
//     Height_cm INT CHECK (height_cm > 100),
//     Strength VARCHAR(50), -- e.g., "Strong", "Moderate"
//     Daily_rate DECIMAL(10,2) CHECK (daily_rate > 0),
//     Experience_years INT CHECK (experience_years >= 0),
//     Profile_picture TEXT, -- URL to the officer's picture
//     Location TEXT,
//     Created_at TIMESTAMP DEFAULT NOW()

type TblSecurityOfficers struct {
	Id              int       `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	UserId          int       `gorm:"column:UserId"`
	DateOfBirth     string    `gorm:"column:DateOfBirth"`
	HeightCm        string    `gorm:"column:HeightCm"`
	Strength        string    `gorm:"column:Strength"`
	DailyRate       string    `gorm:"column:DailyRate"`
	ExperienceYears string    `gorm:"column:ExperienceYears"`
	ProfilePicture  string    `gorm:"column:ProfilePicture"`
	Location        string    `gorm:"column:Location"`
	CreatedAt       time.Time `gorm:"column:CreatedAt"`
	Status          int       `gorm:"column:Status"`
}

type SecurityOfficers struct {
	Id              int       `gorm:"column:Id"`
	UserId          int       `gorm:"column:UserId"`
	DateOfBirth     string    `gorm:"column:DateOfBirth"`
	HeightCm        string    `gorm:"column:HeightCm"`
	Strength        string    `gorm:"column:Strength"`
	DailyRate       string    `gorm:"column:DailyRate"`
	ExperienceYears string    `gorm:"column:ExperienceYears"`
	ProfilePicture  string    `gorm:"column:ProfilePicture"`
	Location        string    `gorm:"column:Location"`
	CreatedAt       time.Time `gorm:"column:CreatedAt"`
	Status          int       `gorm:"column:Status"`
}
