package global

import "github.com/gorilla/sessions"

var (
	Store = sessions.NewCookieStore([]byte("super-secret-key-1234567890abcdef"))
)
