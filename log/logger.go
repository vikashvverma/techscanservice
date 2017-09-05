package log

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"time"
)

type Logger interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type requestResponseLogger struct {
}

func New() Logger {
	return &requestResponseLogger{}
}

func (rrl *requestResponseLogger)  ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	start := time.Now()

	reqFields := requestFields(r)
	logrus.WithFields(reqFields).Infof("Request")

	next(w, r)

	res := w.(negroni.ResponseWriter)
	resFields := responseFields(r, res)
	resFields["Duration"] = int64(time.Since(start) / time.Millisecond)
	logrus.WithFields(resFields).Infof("Request")
}

func requestFields(r *http.Request) logrus.Fields {
	fields := logrus.Fields{}
	fields["Client"] = r.RemoteAddr
	fields["Method"] = r.Method
	fields["URL"] = r.URL.String()
	fields["Referrer"] = r.Referer()
	fields["User-Agent"] = r.UserAgent()
	return fields
}

func responseFields(r *http.Request, w negroni.ResponseWriter) logrus.Fields {
	fields := logrus.Fields{}
	fields["Method"] = r.Method
	fields["URL"] = r.URL.String()
	fields["StatusCode"] = w.Status()
	fields["Size"] = w.Size()
	return fields
}
