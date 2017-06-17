package models

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	Name   string          `json:"name"bson:"name"`
	UserID []bson.ObjectId `json:"user_id"bson:"user_id"`
	ID     bson.ObjectId   `json:"id"bson:"_id"`
	Status bool            `json:"status"bson:"status"`
}

type IssueBook struct {
	UserID string `json:"user_id"`
}

func (b *Book) Create(sess *mgo.Session) error {
	b.ID = bson.NewObjectId()
	if err := sess.DB("library").C("books").Insert(b); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (b *Book) Get(sess *mgo.Session) error {
	if err := sess.DB("library").C("books").FindId(b.ID).One(&b); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (b *Book) Update(sess *mgo.Session) error {
	info, err := sess.DB("library").C("books").UpsertId(b.ID, b)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if info.Updated == 0 {
		return errors.New("Sorry, cannot update user")
	}
	return nil
}

func (b *Book) Delete(sess *mgo.Session) error {
	b.Status = false
	info, err := sess.DB("library").C("books").UpsertId(b.ID, b)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if info.Updated == 0 {
		return errors.New("Sorry, cannot update user")
	}
	return nil
}
