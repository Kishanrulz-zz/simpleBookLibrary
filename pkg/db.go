package db

import "gopkg.in/mgo.v2"

func GetSession(url string) *mgo.Session {
	s, err := mgo.Dial(url) //"mongodb://localhost"
	if err != nil {
		panic(err)
	}
	return s
}
