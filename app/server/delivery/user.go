package delivery

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/moguchev/BD-Forum/pkg/messages"
	. "github.com/moguchev/BD-Forum/pkg/models"

	"github.com/gorilla/mux"
)

func (h *Handler) GetUserByNick(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	code := 200

	nickname, ok := mux.Vars(r)["nickname"]
	if !ok {
		return
	}

	user, err := h.Service.GetUser(nickname)

	var answer []byte

	if err != nil {
		code = 404

		var e Error
		e.Message = messages.UserNotFound + nickname

		answer, _ = json.Marshal(e)
	} else {
		answer, _ = json.Marshal(user)
	}

	w.WriteHeader(code)
	w.Write(answer)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	code := 201

	nickname, ok := mux.Vars(r)["nickname"]
	if !ok {
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	var user User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	user.Nickname = nickname

	users, err := h.Service.CreateUser(user)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	var answer []byte

	if len(users) > 0 {
		code = 409
		answer, _ = json.Marshal(users)
		log.Println(messages.UserAlreadyExists)
	} else {
		answer, _ = json.Marshal(user)
		log.Println(messages.UserWasCreated, user)
	}

	w.WriteHeader(code)
	w.Write(answer)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	code := 200

	nickname, ok := mux.Vars(r)["nickname"]
	if !ok {
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	var user User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	user.Nickname = nickname

	e := h.Service.UpdateUser(user)

	var answer []byte
	var msg Error

	if e != nil {
		if e.Error() == messages.UserNotFound {
			code = 404
			msg.Message = e.Error() + nickname
		}
		if e.Error() == messages.ConflictsInUserUpdate {
			code = 409
			msg.Message = e.Error()
		}
		answer, _ = json.Marshal(msg)
	} else {
		var up UserUpdate
		up.About = user.About
		up.Email = user.Email
		up.Fullname = user.Fullname
		answer, _ = json.Marshal(up)
	}

	w.WriteHeader(code)
	w.Write(answer)
}
