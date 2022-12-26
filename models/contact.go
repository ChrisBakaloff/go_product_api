package models

import (
	"fmt"
	u "github.com/ChrisBakaloff/go_product_api/utils"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "error"), false
	}
	if contact.Phone == "" {
		return u.Message(false, "error"), false
	}
	if contact.UserId <= 0 {
		return u.Message(false, "Error"), false
	}
	return u.Message(true, "success"), true
}

func (contact *Contact) Create() map[string]interface{} {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}
	GetDB().Create(contact)

	resp := u.Message(true, "succcess")
	resp["contact"] = contact
	return resp

}

func GetContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id=?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id=?", user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return contacts
}
