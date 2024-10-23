package middleware

import (
	"net/http"

	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/helpers"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/models"
)

func UserAutheticateJWT(f func(http.ResponseWriter, *http.Request, *models.User)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		user, err := helpers.GetUserFromToken(token)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		f(w, r, user)
	}
}
