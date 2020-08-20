package data

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
)

var once sync.Once

var (
	instance *gorm.DB
)

type UserRepository interface {
	CloseConn()
	GetOrCreate(userResult *User, userWhere *User) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (u *UserRepositoryImpl) CloseConn() {
	u.DB.Close()
}

func (u *UserRepositoryImpl) GetOrCreate(userResult *User, userWhere *User) error {
	return u.DB.FirstOrCreate(userResult, userWhere).Error
}

type User struct {
	ID   string `gorm:"primary_key"`
	Name string
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{DB: initConnection()}
}

func initConnection() *gorm.DB {
	once.Do(func() { // <-- atomic, does not allow repeating
		addr := fmt.Sprintf("postgresql://root@%s:%s/postgres?sslmode=disable", os.Getenv("DB_SERVICE_HOST"), os.Getenv("DB_SERVICE_PORT"))
		db, err := gorm.Open("postgres", addr)
		if err != nil {
			log.Fatal(err)
		}
		db.LogMode(true)
		db.AutoMigrate(&User{})
		instance = db
	})
	return instance
}
