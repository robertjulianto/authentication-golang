package database

import (
	"fmt"
	"ims/types"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	Instance   *gorm.DB
	HOSTNAME   string
	DBUSER     string
	DBPASSWORD string
	DBNAME     string
	PORT       string
	SSLMODE    string
	TIMEZONE   string
}

type Database interface {
	Run()
	GetInstance() *gorm.DB
}

func ConnectToDataBase() (*database, error) {

	envFile, _ := godotenv.Read(".env")

	HOSTNAME := envFile["HOSTNAME"]
	DBUSER := envFile["DBUSER"]
	DBPASSWORD := envFile["DBPASSWORD"]
	DBNAME := envFile["DBNAME"]
	PORT := envFile["PORT"]
	SSLMODE := envFile["SSLMODE"]
	TIMEZONE := envFile["TIMEZONE"]

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		HOSTNAME,
		DBUSER,
		DBPASSWORD,
		DBNAME,
		PORT,
		SSLMODE,
		TIMEZONE,
	)

	instance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return &database{
		Instance:   instance,
		HOSTNAME:   envFile["HOSTNAME"],
		DBUSER:     envFile["DBUSER"],
		DBPASSWORD: envFile["DBPASSWORD"],
		DBNAME:     envFile["DBNAME"],
		PORT:       envFile["PORT"],
		SSLMODE:    envFile["SSLMODE"],
		TIMEZONE:   envFile["TIMEZONE"],
	}, err
}

func (db *database) Run() {
	db.Instance.AutoMigrate(
		types.User{},
		types.Role{},
		types.Access{},
		types.RoleAccess{},
		types.UserRole{},
	)
}

func (db *database) GetInstance() *gorm.DB {
	return db.Instance
}
