package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NOSurf adds CSRF protection to all POST requests
func NOSurf(next http.Handler) http.Handler {
	CSRFHandler := nosurf.New(next)
	CSRFHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return CSRFHandler
}

// SessionLoad loads and saves sessions for every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
