package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/vikashvverma/techscanservice/github"
	"github.com/vikashvverma/techscanservice/response"
)

func Owner(fetcher github.Fetcher, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var repoID int
		var err error
		repoIDString, ok := vars["repoID"]
		if ok {
			repoID, err = strconv.Atoi(repoIDString)
			if err != nil {
				response.ServerError(w)
				return
			}
			if repoID < 0 {
				response.ClientError(w)
				return
			}
		}
		user, err := fetcher.User(int64(repoID))
		if err != nil {
			l.WithError(err).Errorf("Owner: could not get user")
		}
		jsonResponse, err := json.Marshal(user)
		if err != nil {
			e := fmt.Sprintf("error marshalling response: %s", err)

			l.Error(e)
			response.ServerError(w)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(jsonResponse)
	}
}
