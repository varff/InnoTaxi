package auth

import (
	"InnoTaxi/models"
	"InnoTaxi/pkg/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type LoginModel struct {
	Phone int32  `json:"phone"`
	Pass  string `json:"password"`
}

type RegisterModel struct {
	Phone int32  `json:"phone"`
	Pass  string `json:"password"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func CreateToken(phone int32, pass string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["phone"] = phone
	atClaims["pass"] = pass
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(helper.GetEnvDefault("TOKENAC", "notsafe")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func HashPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *gin.Context) {
	var loginMod LoginModel
	if err := c.ShouldBindJSON(&loginMod); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	passHash, err := models.UserLogin(loginMod.Phone)
	if err != nil {
		return
	}
	login := CheckPassHash(loginMod.Pass, passHash)
	if !login {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(loginMod.Phone, loginMod.Pass)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func Register(c *gin.Context) {
	var regModel RegisterModel
	if err := c.ShouldBindJSON(&regModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	passwordHash, err := HashPass(regModel.Pass)
	if err != nil {
		log.Printf("Hashing error: %s", err)
		return
	}
	regged, err := models.UserRegister(regModel.Name, passwordHash, regModel.Email, regModel.Phone)
	if err != nil || !regged {
		c.JSON(http.StatusUnprocessableEntity, "Invalid data")
		return
	}
	token, err := CreateToken(regModel.Phone, regModel.Pass)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
