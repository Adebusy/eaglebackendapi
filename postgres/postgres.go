package postgres

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Adebusy/eaglebackendapi/dataaceess"
	"github.com/Adebusy/eaglebackendapi/obj"

	// "github.com/Adebusy/eaglebackendapi/obj"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DbGorm *gorm.DB
var err error

func GetDB() *gorm.DB {
	// if loadEnv := godotenv.Load(); loadEnv != nil {
	// 	ret := fmt.Sprintf("Unable to load environment variable. %s", loadEnv.Error())
	// 	fmt.Println(ret)
	// }

	env := os.Getenv("env")
	SERVER := os.Getenv("DATABASE_SERVER")
	USERID := os.Getenv("USERID")
	DATABASE := os.Getenv("DATABASE")
	PASSWORD := os.Getenv("PASSWORD")
	PORT := os.Getenv("DB_PORT")

	var dbStatus obj.ConfigStruct
	var connectionString string

	if env == "live" {
		fmt.Println("connected live")
		connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=require", USERID, PASSWORD, SERVER, "25060", DATABASE)
	} else {
		fmt.Println("connected local")
		connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", SERVER, USERID, PASSWORD, DATABASE, PORT)
	}

	logrus.Info(connectionString)
	DbGorm, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, NoLowerCase: true,
	}, PrepareStmt: false})

	if err != nil {
		fmt.Sprintln(err.Error())
		panic("failed to connect database")
	}

	read, err := os.ReadFile("config.json")
	if err != nil {
		logrus.Error(err)
	}
	if err := json.Unmarshal(read, &dbStatus); err != nil {
		logrus.Error(err)
	}

	if dbStatus.CreateTable {
		DbGorm.AutoMigrate(&dataaceess.TblCertificates{})
		DbGorm.AutoMigrate(&dataaceess.TblCourses{})
		DbGorm.AutoMigrate(&dataaceess.TblCoursesSchedule{})
		DbGorm.AutoMigrate(&dataaceess.TblJobs{})
		DbGorm.AutoMigrate(&dataaceess.TblStatus{})
		DbGorm.AutoMigrate(&dataaceess.TblPayments{})
		DbGorm.AutoMigrate(&dataaceess.TblRatings{})
		DbGorm.AutoMigrate(&dataaceess.TblSecurityOfficers{})
		DbGorm.AutoMigrate(&dataaceess.TblSubscriptions{})
		DbGorm.AutoMigrate(&dataaceess.TblTrainingCenters{})
		DbGorm.AutoMigrate(&dataaceess.TblUser{})
	}

	dbStatus.IsDropExistingTables = false
	dbStatus.CreateTable = false
	domarchal, _ := json.Marshal(dbStatus)
	_ = os.WriteFile("config.json", domarchal, 0400)
	return DbGorm
}
