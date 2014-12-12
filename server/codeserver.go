package server

import (
	"encoding/json"
	"github.com/boombuler/hackedinterpreter/token"
	"net/http"
)

type codeServer struct {
	*Server
}

type TokenExport struct {
	token.Pos
	Text string `json:",omitempty"`
	Type string
}

func (cs *codeServer) RegisterHandlers(s *Server, m *http.ServeMux) {
	cs.Server = s
	m.HandleFunc("/code", cs.serveCode)
}

func (cs *codeServer) CodeLoaded() {
	// nothing to do.
}

func (cs *codeServer) serveCode(resp http.ResponseWriter, req *http.Request) {
	if cs.tokens != nil {
		tex := make([]*TokenExport, 0, len(cs.tokens))
		for _, t := range cs.tokens {
			if len(t.Lit) == 0 {
				continue
			}
			x := &TokenExport{
				t.Pos,
				string(t.Lit),
				token.TokMap.Id(t.Type),
			}

			if x.Text == x.Type {
				x.Text = ""
			}
			tex = append(tex, x)
		}
		err := json.NewEncoder(resp).Encode(tex)
		if err != nil {
			panic(err)
		}
	} else {
		resp.WriteHeader(http.StatusNoContent)
	}
}
