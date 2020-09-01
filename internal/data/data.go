package data

import (
	"fmt"
	"log"
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
	GetCharactersByUserId(userId string) ([]UserCharacter, error)
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

func (u *UserRepositoryImpl) GetCharactersByUserId(userId string) ([]UserCharacter, error) {
	var characters []UserCharacter
	err := u.DB.Set("gorm:auto_preload", true).Find(&characters, &UserCharacter{UserId: userId}).Error
	return characters, err
}

func (u *UserRepositoryImpl) AddCharacterToUser(character *Character, userId string) error {
	tx := u.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var user User
	tx.First(&user, "id = ?", userId)
	if user.ID == "" {
		tx.Rollback()
		return fmt.Errorf("Error, no user found for id %s", userId)
	}

	var characterFound Character
	if err := tx.FirstOrCreate(&characterFound, character).Error; err != nil {
		tx.Rollback()
		return err
	}

	var userCharacter UserCharacter
	if err := tx.FirstOrCreate(&userCharacter, &UserCharacter{UserId: user.ID, CharacterId: characterFound.ID}).Error; err != nil {
		tx.Rollback()
		return err
	}

	userCharacter.Count = userCharacter.Count + 1

	if err := tx.Save(&userCharacter).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

type User struct {
	ID   string `gorm:"primary_key"`
	Name string
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

type UserCharacter struct {
	UserId      string `gorm:"primary_key"`
	CharacterId string `gorm:"primary_key"`
	Count       int
	User        User      `gorm:"foreignkey:UserId,PRELOAD:false"`
	Character   Character `gorm:"foreignkey:CharacterId"`
}

func initConnection() *gorm.DB {
	once.Do(func() { // <-- atomic, does not allow repeating
		addr := fmt.Sprintf("postgresql://root@%s:%s/postgres?sslmode=disable", "my-release-cockroachdb-public.default.svc.cluster.local", "26257")
		db, err := gorm.Open("postgres", addr)
		if err != nil {
			log.Fatalln("Error trying to connect to addr:", addr, "err was:", err)
		}
		db.LogMode(true)
		db.AutoMigrate(&Character{})
		db.AutoMigrate(&User{})
		db.AutoMigrate(&UserCharacter{})
		instance = db
	})
	return instance
}
