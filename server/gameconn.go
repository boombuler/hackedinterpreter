package server

import (
	"../lexer"
	"../runtime"
	"code.google.com/p/go.net/websocket"
	"net/http"
	"strconv"
)

type gameServer struct {
	*Server
	keysDown      map[string]bool
	drawcmds      []*DrawCmd
	width, height int
}

func NewGameServer(code runtime.Callable, lex *lexer.DebugLexer, debuggable bool) (*Server, error) {
	gs := &gameServer{
		keysDown: make(map[string]bool),
		drawcmds: make([]*DrawCmd, 0),
		width:    10,
		height:   10,
	}
	s, err := newServer(code, lex, debuggable, func(s *Server, m *http.ServeMux) {
		m.Handle("/", http.FileServer(http.Dir("./content")))
		m.Handle("/game", websocket.Handler(gs.gameHandler))
	})
	if err == nil {
		gs.Server = s
		s.ctx.UI = gs
		return s, nil
	}
	return nil, err
}

type ClientGameCommand struct {
	Command   string `json:"c"`
	Parameter string `json:"p"`
}

type DrawCmd struct {
	Type   string        `json:"c"`
	Params []interface{} `json:"p"`
}

func (s *gameServer) BtnA() (bool, error) {
	down, ok := s.keysDown["a"]
	if !ok {
		return false, nil
	}
	return down, nil
}
func (s *gameServer) BtnB() (bool, error) {
	down, ok := s.keysDown["b"]
	if !ok {
		return false, nil
	}
	return down, nil
}
func (s *gameServer) Down() (bool, error) {
	down, ok := s.keysDown["down"]
	if !ok {
		return false, nil
	}
	return down, nil
}
func (s *gameServer) Up() (bool, error) {
	down, ok := s.keysDown["up"]
	if !ok {
		return false, nil
	}
	return down, nil
}
func (s *gameServer) Right() (bool, error) {
	down, ok := s.keysDown["right"]
	if !ok {
		return false, nil
	}
	return down, nil
}
func (s *gameServer) Left() (bool, error) {
	down, ok := s.keysDown["left"]
	if !ok {
		return false, nil
	}
	return down, nil
}

func (s *gameServer) Width() (int, error) {
	return s.width, nil
}
func (s *gameServer) Height() (int, error) {
	return s.height, nil
}
func (s *gameServer) Draw(x, y int) error {
	s.drawcmds = append(s.drawcmds, &DrawCmd{
		Type:   "draw",
		Params: []interface{}{x, y},
	})
	return nil
}
func (s *gameServer) DrawText(x, y int, text string) error {
	s.drawcmds = append(s.drawcmds, &DrawCmd{
		Type:   "drawText",
		Params: []interface{}{x, y, text},
	})
	return nil
}

func (s *gameServer) gameHandler(ws *websocket.Conn) {
	dcc := make(chan struct{})
	clientCommands := make(chan *ClientGameCommand, 100)

	go func() {
		select {
		case <-s.closeChan:
			close(dcc)
		case <-dcc:
			return
		}
	}()
	go func() { // reader pipe
		defer close(clientCommands)
		for {
			select {
			case <-dcc:
				return
			default:
				cc := new(ClientGameCommand)
				e := websocket.JSON.Receive(ws, cc)
				if e != nil {
					close(dcc)
				} else {
					clientCommands <- cc
				}
			}
		}
	}()

	for {
	input:
		for {
			select {
			case <-dcc:
				return
			case cc, ok := <-clientCommands:
				if ok {
					switch cc.Command {
					case "down":
						s.keysDown[cc.Parameter] = true
					case "up":
						s.keysDown[cc.Parameter] = false
					case "setWidth":
						if w, err := strconv.Atoi(cc.Parameter); err == nil {
							s.width = w
						}
					case "setHeight":
						if h, err := strconv.Atoi(cc.Parameter); err == nil {
							s.height = h
						}
					}
				} else {
					return
				}
			default:
				break input
			}
		}

		s.drawcmds = []*DrawCmd{
			&DrawCmd{Type: "clear", Params: []interface{}{}},
		}
		_, err := s.ctx.Call(s.code)
		if err != nil {
			websocket.JSON.Send(ws, []*DrawCmd{
				&DrawCmd{
					Type:   "error",
					Params: []interface{}{err.Error()},
				},
			})
		}
		websocket.JSON.Send(ws, s.drawcmds)
	}
}
