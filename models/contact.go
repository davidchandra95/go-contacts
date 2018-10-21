package models

import (
	"log"
	"github.com/jinzhu/gorm"
	u "go-contacts/utils"
	"fmt"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	User_FK uint   `json:"user_id"` //The user FK
}

func (contact *Contact) Validate() (map[string] interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Contact name is required."), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number is required."), false
	}

	if contact.User_FK <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "Passed validation"), true
}

func (contact *Contact) Create() (map[string] interface{}) {
	log.Println("we are in create func")
	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "Success adding contact")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) (*Contact) {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user_FK uint) ([]*Contact) {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("User_FK = ?", user_FK).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return contacts
}