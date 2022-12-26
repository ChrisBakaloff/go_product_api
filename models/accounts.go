package models

import (
	u "github.com/ChrisBakaloff/go_product_api/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func GetUser(u uint) *Account {
	acc := &Account{}
	GetDB().Table("accounts").Where("id=?", u).First(acc)
	if acc.Email == "" {
		return nil
	}
	acc.Password = ""
	return acc
}

func Login(email, password string) (map[string]interface{}) {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login creds")
	}

	account.Password = ""
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	resp := u.Message(true, "Logged in")
	resp["account"] = account
	return resp
}

func (account *Account) Create() (map[string]interface{}) {
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	//Generate the password hash using bcrypt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	GetDB().Create(account)

	//Check if account ID has been created
	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, error")
	}

	//Create JWT token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = ""

	resp := u.Message(true, "Account has been created")
	resp["account"] = account
	return resp
}

func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password length not sufficient"), false
	}

	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error"), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email address already in use"), false
	}

	return u.Message(false, "Requirement passed"), true
}
