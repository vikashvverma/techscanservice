package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/vikashvverma/techscanservice/github"
	"github.com/vikashvverma/techscanservice/response"
	"strconv"
	"strings"
)

func Language(fetcher github.Fetcher, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		language := vars["lang"]

		page := 0
		var err error
		pageString, ok := vars["page"]
		if ok {
			page, err = strconv.Atoi(pageString)
			if err != nil {
				response.ServerError(w)
				return
			}
			if page < 0 {
				response.ClientError(w)
				return
			}
		}

		repositories, err := fetcher.Language(strings.ToLower(language), int64(page))

		if err != nil {
			l.WithError(err).Errorf("Language: could not repositories")
		}
		jsonResponse, err := json.Marshal(repositories)
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
