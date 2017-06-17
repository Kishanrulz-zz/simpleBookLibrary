package controllers

import (
	"github.com/kishanrulz/simpleBookLibrary/model"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"stash.bms.bz/mice/media/errors"
)

func (api Api) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", "Sorry, invalid input")
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", "Sorry, invalid input")
		return
	}

	if err := u.Create(api.session); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", uj)
		return
	} else {
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "%s\n", uj)
		return 
	}
}

// Methods have to be capitalized to be exported, eg, GetUser and not getUser
func (api Api) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := &models.User{
		ID: bson.ObjectIdHex(p.ByName("id")),
	}

	if err := user.Get(api.session); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", user)
	}

	uj, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%s\n", uj)

}

func (api Api) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", user)
	}
	user.ID = bson.ObjectIdHex(p.ByName("id"))

	if err := user.Update(api.session); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", user)
	}

	uj, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%s\n", uj)
}

func (api Api) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", user)
	}
	user.ID = bson.ObjectIdHex(p.ByName("id"))

	if err := user.Delete(api.session); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", user)
	}

	uj, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%s\n", uj)
}

func isValid(userID bson.ObjectId, sess *mgo.Session) error {
	user := &models.User{
		ID: userID,
	}
	if err := user.Get(sess); err != nil {
		return err
	}
	if user.Status == false {
		return errors.New("Invalid user")
	}
	return nil
}
