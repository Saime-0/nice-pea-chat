package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"gorm.io/gorm"

	serviceL10n "github.com/saime-0/nice-pea-chat/internal/service/l10n"
)

type ServerParams struct {
	Ctx  context.Context
	Addr string

	L10n    serviceL10n.Service
	DB      *gorm.DB
	PGXconn *pgx.Conn
}

func (s ServerParams) StartServer() error {
	muxHttp := http.NewServeMux()
	s.declareRoutes(muxHttp)
	serverHttp := &http.Server{
		Addr:        s.Addr,
		Handler:     muxHttp,
		ReadTimeout: 5 * time.Minute,
	}

	go func() {
		<-s.Ctx.Done()
		if err := serverHttp.Shutdown(context.Background()); err != nil {
			log.Printf("[StartServer] closed with error: %v", err)
		}
	}()

	err := serverHttp.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
