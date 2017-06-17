package controllers

import (
	"ExamplesDemo/Lib/pkg"
	"gopkg.in/mgo.v2"
)

type Api struct {
	session *mgo.Session
}

func NewController(url string) *Api {
	return &Api{db.GetSession(url)}
}
