package cookie

import (
	"net/http"
	"time"
)

func SetCookie(name, value string, w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:    name,
		Value:   value,
		Path:    "/",
		Expires: time.Now().Add(365 * 24 * time.Hour),
	}

	http.SetCookie(w, cookie)
}
