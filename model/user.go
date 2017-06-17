package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"stash.bms.bz/mice/media/errors"
)

type User struct {
	Name   string        `json:"name"bson:"name"`
	Gender string        `json:"gender"bson:"gender"`
	Age    int           `json:"age"bson:"age"`
	ID     bson.ObjectId `json:"id"bson:"_id"`
	Status bool          `json:"status"bson:"status"`
}

func (u *User) Create(sess *mgo.Session) error {
	u.ID = bson.NewObjectId()
	if err := sess.DB("library").C("users").Insert(u); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *User) Get(sess *mgo.Session) error {
	if err := sess.DB("library").C("users").FindId(u.ID).One(&u); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *User) Update(sess *mgo.Session) error {
	info, err := sess.DB("library").C("users").UpsertId(u.ID, u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if info.Updated == 0 {
		return errors.New("Sorry, cannot update user")
	}
	return nil
}

func (u *User) Delete(sess *mgo.Session) error {
	u.Status = false
	info, err := sess.DB("library").C("users").UpsertId(u.ID, u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if info.Updated == 0 {
		return errors.New("Sorry, cannot update user")
	}
	return nil
}
