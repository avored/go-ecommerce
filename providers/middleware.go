package providers

import (
	"encoding/base32"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/securecookie"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func MultipleMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {

	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped

}



func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := GetSessionStore(r)
		if auth, ok := session.Values["AdminUser_Authenticated"].(bool); !ok || !auth {
			NewUrl := "/admin/login"
			http.Redirect(w, r, NewUrl, http.StatusSeeOther)
			return
		}
		h(w, r)
	}
}
func CustomerAuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := GetSessionStore(r)
		if auth, ok := session.Values["Customer_Authenticated"].(bool); !ok || !auth {
			NewUrl := "/login"
			http.Redirect(w, r, NewUrl, http.StatusSeeOther)
			return
		}
		h(w, r)
	}
}

func SessionMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := GetSessionStore(r)
		if session.Values["ID"] == nil {
			fmt.Println("******** SESSION ID GENERATING GOING ON **********")
			// Generate a random session ID key suitable for storage in the DB
			session.Values["ID"] = string(securecookie.GenerateRandomKey(32))
			session.Values["ID"] = strings.TrimRight(
				base32.StdEncoding.EncodeToString(
					securecookie.GenerateRandomKey(32)), "=")
		}

		session.Save(r, w)

		fmt.Printf("SESSION ID: %v", session.Values["ID"])
		h(w, r)
	}
}
