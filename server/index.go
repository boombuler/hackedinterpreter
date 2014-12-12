package server

import (
	"net/http"
)

type indexMod struct {
	*Server
}

func (im *indexMod) RegisterHandlers(s *Server, m *http.ServeMux) {
	im.Server = s
	m.HandleFunc("/index", im.redirect)
}

func (im *indexMod) CodeLoaded() {
	// nothing to do.
}

func (im *indexMod) redirect(w http.ResponseWriter, r *http.Request) {
	if im.code == nil {
		http.Redirect(w, r, "/loadcode.html", http.StatusTemporaryRedirect)
	}
}
