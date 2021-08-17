package service

import (
	"log"
	"portfolio/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct{}

func (u User) Me(username string) (me model.User, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Select("username, is_active, is_superuser").Table("auth_user").Where("username = ?", username).Find(&me).Error; err != nil {
		return
	}

	return
}

func (u User) GetUsers() (users []model.User, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Select("username, is_active, is_superuser").Table("auth_user").Find(&users).Error; err != nil {
		return
	}

	return
}

func (u User) CreateUser(username, password string) (err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	_pw := []byte(password)
	hashed, _ := bcrypt.GenerateFromPassword(_pw, 10)
	if err = bcrypt.CompareHashAndPassword(hashed, _pw); err != nil {
		return
	}

	t := time.Now()
	user := model.User{Username: username, Password: hashed, IsActive: true, IsStaff: true, IsSuperuser: true, DateJoined: t}
	if err := db.Table("auth_user").Create(&user).Error; err != nil {
		return err
	}

	return
}

// 妥協
// あっちで作るとライフサイクル的にキツかった
// import cycle not allowed
// package main
// 	imports portfolio/router
// 	imports portfolio/controller
// 	imports portfolio/service
// 	imports portfolio/middleware
// 	imports portfolio/service
func createToken(username string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256")) // JWT
	token.Claims = jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	var secretKey = "takurinton"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Login(username, password string) (token string, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	var u model.User
	if err = db.Table("auth_user").Select("password").Where("username = ?", username).Find(&u).Error; err != nil {
		return
	}

	// CompareHashAndPasswordの引数がbyte型だったからGenerateして渡してたけどそうじゃなかったっぽい
	// _pw := []byte(password)
	// hashed, _ := bcrypt.GenerateFromPassword(_pw, 10)
	// fmt.Println(hashed)
	if err = bcrypt.CompareHashAndPassword(u.Password, []byte(password)); err != nil {
		log.Println(err)
		return
	}

	token, err = createToken(username)
	if err != nil {
		return
	}
	return token, nil
}
