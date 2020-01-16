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

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	code := 201

	nickname, ok := mux.Vars(r)["nickname"]
	if !ok {
		//
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//
	}

	var newUser NewUser
	err = json.Unmarshal(bytes, &newUser)
	if err != nil {

	}

	user, err := h.Service.CreateUser(newUser, nickname)
	if err != nil {
		if err.Error() == messages.UserAlreadyExists {
			code = 409
		}
		// if err.Error() == messages.UserNotFound {
		// 	SetError(w, 404, err.Error())
		// 	return
		// }
		log.Println(err.Error())
	}

	answer, _ := json.Marshal(user)

	w.WriteHeader(code)
	w.Write(answer)
}
