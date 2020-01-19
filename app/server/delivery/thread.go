package delivery

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moguchev/BD-Forum/pkg/messages"
	. "github.com/moguchev/BD-Forum/pkg/models"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	code := 201

	slugOrID, ok := mux.Vars(r)["slug_or_id"]
	if !ok {
		return
	}

	var posts []Post

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	err = json.Unmarshal(bytes, &posts)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	posts, err = h.Service.CreatePosts(slugOrID, posts)

	answer, _ := json.Marshal(posts)

	if err != nil {
		switch err.Error() {
		case messages.ThreadNotFound:
			code = 404
			answer, _ = json.Marshal(Error{Message: err.Error() + slugOrID})
			break
		case messages.ParentNotFound:
			code = 409
			answer, _ = json.Marshal(Error{Message: err.Error()})
			break
		default:
			code = 409
			answer, _ = json.Marshal(Error{Message: err.Error()})
		}
	}

	w.WriteHeader(code)
	w.Write(answer)
}
