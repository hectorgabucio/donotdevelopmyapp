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

// USERS //////////////////////////////////////////////////////////////////////////////

type UserRepository interface {
	CloseConn()
	GetOrCreate(userResult *User, userWhere *User) error
	AddCharacterToUser(character *Character, userId string) error
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

func (u *UserRepositoryImpl) AddCharacterToUser(character *Character, userId string) error {
	var user User
	u.DB.First(&user, "id = ?", userId)
	if user.ID == "" {
		return fmt.Errorf("Error, no user found for id %s", userId)
	}

	var characterFound Character
	if err := u.DB.FirstOrCreate(&characterFound, character).Error; err != nil {
		return err
	}

	return u.DB.Model(&user).Association("Characters").Append(characterFound).Error

}

type User struct {
	ID         string `gorm:"primary_key"`
	Name       string
	Characters []Character `gorm:"many2many:user_characters;association_foreignkey:id;foreignkey:id"`
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{DB: initConnection()}
}

// CHARACTERS FROM USERS ////////////////////////////////////////////////////////////////////

type Character struct {
	ID    string `gorm:"primary_key"`
	Name  string
	Image string
}

func initConnection() *gorm.DB {
	once.Do(func() { // <-- atomic, does not allow repeating
		addr := fmt.Sprintf("postgresql://root@%s:%s/postgres?sslmode=disable", os.Getenv("DB_SERVICE_HOST"), os.Getenv("DB_SERVICE_PORT"))
		db, err := gorm.Open("postgres", addr)
		if err != nil {
			log.Fatal(err)
		}
		db.LogMode(true)
		db.AutoMigrate(&Character{})
		db.AutoMigrate(&User{})
		instance = db
	})
	return instance
}
