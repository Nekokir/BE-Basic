package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"time"
	"strconv"
)

type User struct {
	ID        uint `gorm:"AUTO_INCREMENT"`
	Name      string
	Token     string
	Answers   []Answer
	Submitted bool
}

type Answer struct {
	ID     uint `gorm:"AUTO_INCREMENT"`
	QID    uint
	UserID uint
	Answer string
}

func FindUser(db *gorm.DB, token string) (user User, err error) {
	err = db.Where("token = ?", token).Where("submitted = ?", false).First(&user).Error
	if err != nil {
		err = errors.Wrap(err, "FindUser")
		return
	}
	return
}

func FindSubmittedUsers(db *gorm.DB) (users []User, err error)  {
	err = db.Where("submitted = ?", true).Find(&users).Error
	if err != nil {
		err = errors.Wrap(err, "FindSubmittedUsers")
		return
	}
	return
}

func FindDetailedUsers(db *gorm.DB) (users []User, err error)  {
	err = db.Preload("Answers").Order("submitted desc").Find(&users).Error
	if err != nil {
		err = errors.Wrap(err, "FindDetailedUsers")
		return
	}
	return
}

func StoreAnswer(db *gorm.DB, uid uint, qid uint, ans string) (answer Answer, err error) {
	var count int
	db.Model(&Answer{}).Where("user_id = ?", uid).Where("q_id = ?", qid).Count(&count)
	if count != 0 {
		err = errors.New("Answer already exist")
		return
	}
	answer.UserID = uid
	answer.QID = qid
	answer.Answer = ans
	err = db.Create(&answer).Error
	if err != nil {
		err = errors.Wrap(err, "StoreAnswer")
		return
	}
	return
}

func FindAnswer(db *gorm.DB, uid uint, qid uint) (answer Answer, err error) {
	err = db.Where("user_id = ?", uid).Where("q_id = ?", qid).First(&answer).Error
	if err != nil {
		err = errors.Wrap(err, "FindAnswer")
		return
	}
	return
}

func UpdateAnswer(db *gorm.DB, uid uint, qid uint, ans string) (answer Answer, err error) {
	err = db.Where("user_id = ?", uid).Where("q_id = ?", qid).First(&answer).Error
	if err != nil {
		err = errors.Wrap(err, "UpdateAnswer")
		return
	}
	answer.Answer = ans
	err = db.Model(&answer).Updates(answer).Error
	if err != nil {
		err = errors.Wrap(err, "UpdateAnswer")
		return
	}
	return
}

func StoreUser(db *gorm.DB, name string) (user User, err error)  {
	user.Name = name
	user.Token = GetMD5Hash(name + strconv.FormatInt(time.Now().Unix(), 10))
	user.Submitted = false
	err = db.Create(&user).Error
	if err != nil {
		err = errors.Wrap(err, "StoreUser")
		return
	}
	return
}

func UpdateUser(db *gorm.DB, token string, name string) (user User, err error)  {
	err = db.Preload("Answers").Where("token = ?", token).First(&user).Error
	if err != nil {
		err = errors.Wrap(err, "UpdateUser")
		return
	}
	user.Name = name
	err = db.Model(&user).Updates(user).Error
	if err != nil {
		err = errors.Wrap(err, "UpdateUser")
		return
	}
	return
}

func SubmitUser(db *gorm.DB, token string) (user User, err error)  {
	err = db.Preload("Answers").Where("token = ?", token).First(&user).Error
	if err != nil {
		err = errors.Wrap(err, "SubmitUser")
		return
	}
	user.Submitted = true
	user.Token = "expired"
	err = db.Model(&user).Updates(user).Error
	if err != nil {
		err = errors.Wrap(err, "SubmitUser")
		return
	}
	return
}