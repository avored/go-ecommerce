package providers

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore
type BaseFlashes struct {
	Messages	[]string
	Errors		[]string
	Warnings	[]string
}


func SetSession () {
	sessionKey := os.Getenv("SESSION_KEY")
	key := []byte(sessionKey)
    store = sessions.NewCookieStore(key)

	
}


// Set adds a new message into the cookie storage.
func Set(w http.ResponseWriter, r *http.Request, name, value string) {
	session, _ := GetSessionStore(r)
	session.AddFlash(value, name)

	session.Save(r, w)
}

// Get gets flash messages from the cookie storage.
func Get(w http.ResponseWriter, r *http.Request, name string) []string {
	session, _ := GetSessionStore(r)
	fm := session.Flashes(name)
	// If we have some messages.
	if len(fm) > 0 {
		session.Save(r, w)
		// Initiate a strings slice to return messages.
		var flashes []string
		for _, fl := range fm {
			// Add message to the slice.
			flashes = append(flashes, fl.(string))
		}

		return flashes
	}
	return nil
}


func GetSessionStore(r *http.Request) (*sessions.Session, error) {
	sessionKey := os.Getenv("SESSION_KEY")
	return store.Get(r, sessionKey)
}
