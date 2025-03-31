package dataaceess

import "time"

// CREATE TABLE TblCertificates (
//     Certificate_id SERIAL PRIMARY KEY,
//     Officer_id INT REFERENCES security_officers(officer_id) ON DELETE CASCADE,
//     Training_center VARCHAR(255) NOT NULL,
//     Certification_name VARCHAR(255) NOT NULL,
//     Issued_date DATE NOT NULL,
//     Expiry_date DATE,
//     Certificate_url TEXT, -- URL to view certificate proof
//     UNIQUE (officer_id, certification_name)
// );

type TblCertificates struct {
	Id                int       `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	OfficerId         int       `gorm:"column:OfficerId"`
	CourseId          int       `gorm:"column:CourseId"`
	TrainingCenter    string    `gorm:"column:TrainingCenter"`
	CertificationName string    `gorm:"column:CertificationName"`
	IssuedDate        string    `gorm:"column:IssuedDate"`
	ExpiryDate        string    `gorm:"column:ExpiryDate"`
	CertificateUrl    string    `gorm:"column:CertificateUrl"`
	ProfilePicture    string    `gorm:"column:ProfilePicture"`
	Location          string    `gorm:"column:Location"`
	CreatedAt         time.Time `gorm:"column:CreatedAt"`
	Status            int       `gorm:"column:Status"`
}

type Certificates struct {
	Id                int       `gorm:"column:Id"`
	OfficerId         int       `gorm:"column:OfficerId"`
	CourseId          int       `gorm:"column:CourseId"`
	TrainingCenter    string    `gorm:"column:TrainingCenter"`
	CertificationName string    `gorm:"column:CertificationName"`
	IssuedDate        string    `gorm:"column:IssuedDate"`
	ExpiryDate        string    `gorm:"column:ExpiryDate"`
	CertificateUrl    string    `gorm:"column:CertificateUrl"`
	ProfilePicture    string    `gorm:"column:ProfilePicture"`
	Location          string    `gorm:"column:Location"`
	CreatedAt         time.Time `gorm:"column:CreatedAt"`
	Status            int       `gorm:"column:Status"`
}
