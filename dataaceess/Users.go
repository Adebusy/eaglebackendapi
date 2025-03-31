package dataaceess

type TblUserType struct {
	Id        int    `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	Name      string `gorm:"column:Name"`
	Status    string `gorm:"column:Status"`
	CreatedAt string `gorm:"column:CreatedAt"`
}

type User struct {
	Id           int    `gorm:"column:Id"`
	TitleId      string `gorm:"column:TitleId"`
	UserName     string `gorm:"column:UserName"`
	NickName     string `gorm:"column:NickName"`
	FirstName    string `gorm:"column:FirstName"`
	LastName     string `gorm:"column:LastName"`
	EmailAddress string `gorm:"column:EmailAddress"`
	MobileNumber string `gorm:"column:MobileNumber"`
	Gender       string `gorm:"column:Gender"`
	UserTypeId   int    `gorm:"column:UserType"`
	Address      string `gorm:"column:Address"`
	Location     string `gorm:"column:Location"`
	AgeRange     string `gorm:"column:AgeRange"`
	Password     string `gorm:"column:Password"`
	Status       string `gorm:"column:Status"`
	CreatedAt    string `gorm:"column:CreatedAt"`
}

type TblUser struct {
	Id           int    `json:"Id" gorm:"unique;primaryKey;autoIncrement"`
	TitleId      string `gorm:"column:TitleId"`
	UserName     string `gorm:"column:UserName"`
	NickName     string `gorm:"column:NickName"`
	FirstName    string `gorm:"column:FirstName"`
	LastName     string `gorm:"column:LastName"`
	EmailAddress string `gorm:"column:EmailAddress"`
	MobileNumber string `gorm:"column:MobileNumber"`
	Gender       string `gorm:"column:Gender"`
	UserTypeId   int    `gorm:"column:UserType"`
	Address      string `gorm:"column:Address"`
	Location     string `gorm:"column:Location"`
	AgeRange     string `gorm:"column:AgeRange"`
	Password     string `gorm:"column:Password"`
	Status       string `gorm:"column:Status"`
	CreatedAt    string `gorm:"column:CreatedAt"`
}

type ResponseMessage struct {
	ResponseCode    string
	ResponseMessage string
}
