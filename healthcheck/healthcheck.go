package healthcheck

import (
	"io"
	"net/http"
)

func Self(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I am alive")
}
