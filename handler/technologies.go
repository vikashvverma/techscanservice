package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"

	"github.com/vikashvverma/techscanservice/github"
	"github.com/vikashvverma/techscanservice/response"
)

func Technology(fetcher github.Fetcher, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		languages, err := fetcher.Fetch()
		if err != nil {
			l.WithError(err).Errorf("Technology: could not fetch all languages' rep")
		}
		jsonResponse, err := json.Marshal(languages)
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
