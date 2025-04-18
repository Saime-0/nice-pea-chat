package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5"
	"gorm.io/gorm"

	"github.com/saime-0/nice-pea-chat/internal/service/l10n"
)

type mux struct {
	*http.ServeMux
	s ServerParams
}

type Request struct {
	*http.Request
	L10n    l10n.Service
	DB      *gorm.DB
	PGXconn *pgx.Conn
	Locale  string
	Token   string
}

type HandlerFunc func(Request) (any, error)

func (m *mux) handle(pattern string, f HandlerFunc) {
	m.Handle(pattern, modulation(initRequest(m.s, f)))
}

func modulation(next HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			data any
			b    []byte
			err  error
		)
		w.Header().Add("Content-Type", "application/json")
		if err = r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			if data, err = next(Request{Request: r}); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				data = ResponseError{Error: err.Error()}
			}
			if s, ok := data.(string); ok {
				data = ResponseMsg{Message: s}
			}
		}
		// Marshal data
		if b, err = json.Marshal(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Del("Content-Type")
			w.Header().Add("Content-Type", "text/plain")
			b = []byte(err.Error())
		}
		// Try to send data
		if _, err = w.Write(b); err != nil {
			log.Println("[modulation] error response write:", err.Error())
			return
		}
	}
}

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseMsg struct {
	Message string `json:"message"`
}

func initRequest(s ServerParams, next HandlerFunc) HandlerFunc {
	return func(r Request) (any, error) {
		r = Request{
			Request: r.Request,
			L10n:    s.L10n,
			Locale:  locale(r.Header.Get("Accept-Language"), l10n.LocaleDefault),
			Token:   getToken(r.Request),
			DB:      s.DB,
			PGXconn: s.PGXconn,
		}
		return next(r)
	}
}

const authzHeader = "Authorization"
const authzTypeUser = "NpcUserToken"

func getToken(r *http.Request) string {
	header := r.Header.Get(authzHeader)
	token, _ := strings.CutPrefix(header, authzTypeUser+" ")
	return token
}
